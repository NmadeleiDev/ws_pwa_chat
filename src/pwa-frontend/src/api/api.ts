import axios from 'axios'

interface httpPayload {
    data: Object | null
    auth: Object | null
}

function stringifyParams(params: Object) {
    // @ts-ignore
    return Object.keys(params).map(key => key + '=' + params[key]).join('&');
}

function getRequestUri(path: string, isFile: boolean) {
    if (isFile)
        return '/api/media/' + path;
    else
        return '/api/v1/' + path;
}

const get = async function (path: string, params: Object | null = null) {
    let uri = getRequestUri(path, false);
    if (params)
        uri += ("?" + stringifyParams(params))
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
    const uri = getRequestUri(path, false);
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
    const uri = getRequestUri(path, false);
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

const upload = async function (path: string, payload: FormData) {
    const uri = getRequestUri(path, true);
    let response;
    try {
        response = await axios.post(uri, payload, {headers: {'Content-Type': 'multipart/form-data', 'mobile': 'true'}});
    } catch (e) {
        console.log("Request put error: ", e)
        return {status: false, data: null}
    }
    if (response.data.status !== true) {
        console.log("failed response: ", response);
    }

    return response.data;
};

const download = async function (path: string, params: Object | null = null) {
    let uri = getRequestUri(path, true);
    if (params)
        uri += ("?" + stringifyParams(params))
    let response;
    try {
        response = await axios.get(uri, {headers: {mobile: 'true'}, responseType: 'blob'});
    } catch (e) {
        console.log("Request get error: ", e)
        return {status: false, data: null}
    }
    console.log("Got response: ", response)

    return response.data;
};

export default {
    get,
    put,
    post,
    upload,
    download
};
