// @ts-ignore
import {sha256, sha224} from 'js-sha256';

const generateUserSecret = (user: {username: string, password: string}) => {
    return sha256([user.username, user.password, user.username].join("&-")).toString();
}

const generateToken = (payload: { username: string, timeStamp: string, sessionKey: string, userSecret: string }) => {
    return sha256([payload.username, payload.timeStamp, payload.sessionKey, payload.userSecret].join("")).toString();
}

const getSha224 = (raw: string): string => {
    return sha224(raw)
}

export default {
    generateToken,
    generateUserSecret,
    getSha224
}