import axios from 'axios'

interface httpPayload {
    data: Object | null
    auth: Object | null
}

function getRequestUri(path: string) {
    return '/api/v1/' + path;
}

const get = async function (path: string) {
    const uri = getRequestUri(path);
    let response;
    try {
        response = await axios.get(uri, {headers: {mobile: 'true'}});
    } catch (e) {
        console.log("Request get error: ", e)
        return {status: false, data: null}
    }
    const resultStatus = response.data.status;

    if (resultStatus !== true) {
        console.log(response);
    }

    return response.data;
};

const post = async function (path: string, payload: httpPayload | null, timestamp: string) {
    const uri = getRequestUri(path);
    let response;
    try {
        response = await axios.post(uri, payload, {headers: {'mobile': 'true', 'Event-date': timestamp} });
    } catch (e) {
        console.log("Request post error: ", e)
        return {status: false, data: null}
    }
    if (response.data.status !== true) {
        console.log("failed response: ", response);
    }

    return response.data;
};

const put = async function (path: string, payload: httpPayload | null, timestamp: string) {
    const uri = getRequestUri(path);
    let response;
    try {
        response = await axios.put(uri, payload, {headers: {'mobile': 'true', 'Event-date': timestamp}});
    } catch (e) {
        console.log("Request put error: ", e)
        return {status: false, data: null}
    }
    if (response.data.status !== true) {
        console.log("failed response: ", response);
    }

    return response.data;
};

export default {
    get,
    put,
    post,
};
