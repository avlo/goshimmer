import {action, observable} from 'mobx';
import {registerHandler, WSMsgType} from "app/misc/WS";

class ManaMsg {
    nodeID: string;
    access: number;
    consensus: number;
    // in ms?
    time: string;
}

const maxRegistrySize = 100;

export class ManaStore {
    // mana values
    @observable manaValues: Array<any> = [];

    constructor() {
        this.manaValues = [];
        registerHandler(WSMsgType.Mana, this.addNewManaValue);
    };

    @action
    addNewManaValue = (manaMsg: ManaMsg) =>  {
        if (this.manaValues.length === maxRegistrySize) {
            // shift if we already have enough values
            this.manaValues.shift();
        }
        // TODO: date parsing is not right
        let newManaData = [new Date(manaMsg.time), manaMsg.access, manaMsg.consensus];
        this.manaValues.push(newManaData);
    }
}

export default ManaStore;