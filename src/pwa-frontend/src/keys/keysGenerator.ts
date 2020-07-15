// @ts-ignore
import {sha256} from 'js-sha256';

const generateUserSecret = (user: {username: string, password: string}) => {
    return sha256([user.username, user.password, user.username].join("&-")).toString();
}

const generateToken = (payload: { username: string, timeStamp: string, sessionKey: string, userSecret: string }) => {
    return sha256([payload.username, payload.timeStamp, payload.sessionKey, payload.userSecret].join("")).toString();
}

export default {
    generateToken,
    generateUserSecret
}