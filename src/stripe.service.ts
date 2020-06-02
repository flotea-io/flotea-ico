/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import axios from 'axios';
declare var jGrowl: any;
declare var $: any;

const api = '/stripe';

export default class StripeService {

	stripe: any;
	card: any;
	main: any;

	init(main: any){
		this.main = main;
		this.stripe = Stripe('pk_test_FG1xyNruCNcwSqyk8cFx2FzW005M3TV6uR');

		let elements = this.stripe.elements();

		let style = {
			invalid: {
				color: '#fa755a',
				iconColor: '#fa755a'
			}
		};

		this.card = elements.create('card', {style: style});
		this.card.mount($('#main #card-element')[0]);

		this.card.addEventListener('change', (event: any) => this.change(event));
		$('#main #payment-form')[0].addEventListener('submit', (event: any) => this.submit(event));
	}

	change(event: any){
		$('#main #card-errors').text(event.error? event.error.message : '');
	}

	submit(event: any){
		event.preventDefault();
		if(this.main.customerAddress == "" || !this.main.customerValid){
			if(this.main.customerAddress == "")
				this.main.customerAddress = "0x";
			$.jGrowl("Please fill your ETH address", { header: 'Error', life: 20000 });
		}
		else {
			this.stripe.createToken(this.card).then((result: any) => {
				if (result.error) {
					$('#main #card-errors').text(result.error.message);
				} else {
					this.checkout(result.token.id);
					console.log(result);
				}
			});
		}
	}

	checkout(token: string) {
		let qs = require('qs');
		axios.post(`${api}/buy`, qs.stringify({
			tokens: (parseFloat(this.main.statusAmount)*100).toString(),
			source: token,
			address: this.main.customerAddress
		}), {
			headers: {
    			'Content-Type': 'application/x-www-form-urlencoded'
  			}
  		})
		.then(function (response) {
			console.log(response);
			if(response.data.error){
				if(response.data.data.indexOf("{") == 0){
					$.jGrowl(JSON.parse(response.data.data).message, { header: 'Error', life: 20000 });
					console.log(JSON.parse(response.data.data), response.data.data);
				} else {
					$.jGrowl(response.data.data, { header: 'Error', life: 20000 });
				}
			} else {
				$.jGrowl("You bought tokens, You can refresh token status on the right of address input.", { header: 'Congratulation', life: 20000 });
			}
		})
		.catch(function (error) {
			console.log(error);
		});
	}
}
export const stripeService = new StripeService();
