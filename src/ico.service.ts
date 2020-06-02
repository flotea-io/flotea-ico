/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import axios from 'axios';

const api = '/ico';

export default class IcoService {
  constructor() {
  }


  getInfo() {
    return axios.get<any>(`${api}/info`);
  }

  getTotalPrice(tokens: any) {
    return axios.get<any>(`${api}/price/${tokens}`);
  }

  getTokensForAmount(tokens: any) {
    return axios.get<any>(`${api}/tokens/${tokens}`);
  }

  getEthPrice(){
    return axios.get<any>(`${api}/ethprice`);
  }

  checkAddress(address : any){
    return axios.get<any>(`${api}/checkaddress/${address}`);
  }

  getAddress(){
    return axios.get<any>(`${api}/address`);
  }

  getBalancePln(){
    return axios.get<any>(`${api}/balancepln`);
  }

}
export const icoService = new IcoService();
