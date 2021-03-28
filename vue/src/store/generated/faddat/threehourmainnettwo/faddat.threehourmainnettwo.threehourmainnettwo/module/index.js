// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreatePost } from "./types/threehourmainnettwo/tx";
import { MsgDeletePost } from "./types/threehourmainnettwo/tx";
import { MsgUpdatePost } from "./types/threehourmainnettwo/tx";
const types = [
    ["/faddat.threehourmainnettwo.threehourmainnettwo.MsgCreatePost", MsgCreatePost],
    ["/faddat.threehourmainnettwo.threehourmainnettwo.MsgDeletePost", MsgDeletePost],
    ["/faddat.threehourmainnettwo.threehourmainnettwo.MsgUpdatePost", MsgUpdatePost],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreatePost: (data) => ({ typeUrl: "/faddat.threehourmainnettwo.threehourmainnettwo.MsgCreatePost", value: data }),
        msgDeletePost: (data) => ({ typeUrl: "/faddat.threehourmainnettwo.threehourmainnettwo.MsgDeletePost", value: data }),
        msgUpdatePost: (data) => ({ typeUrl: "/faddat.threehourmainnettwo.threehourmainnettwo.MsgUpdatePost", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
