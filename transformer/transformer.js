import { ethers } from "ethers"


export const process = async (data) => {
    
    if (data.Id !== undefined) {
       data.Id = "changed_id"
    }

    if (data.Number !== undefined) {
        data.Number = data.Number * Math.random() *1000
    }

    /*const delay = ms => new Promise(resolve => setTimeout(resolve, ms))
    await delay(10000)*/
    
    
    return data
}

