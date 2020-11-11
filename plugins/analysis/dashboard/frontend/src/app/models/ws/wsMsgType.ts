export enum WSMsgType {
    ping = 0,
    fpc = 1,
    addNode = 2,
    removeNode = 3,
    connectNodes = 4,
    disconnectNodes = 5,
    Mana= 8,
    ManaMapOverall,
    ManaMapOnline,
    ManaAllowedPledge,
    ManaPledge,
    ManaRevoke,
    MsgManaDashboardAddress,
    MsgReqManaDashboardAddress
}