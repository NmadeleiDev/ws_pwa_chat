import axios from 'axios'
import { alertNotifierChannel } from '../main';

const get = async function (path) {
    const uri = getRequestUri(path);
    let resultData;
    let resultStatus;

    let response = await axios.get(uri);
    resultData = response.data.data;
    resultStatus = response.data.status;
    console.log("GET: resultStatus: ", resultStatus);

    if (resultStatus !== true) {
        console.log(response);
    }

    return response.data;
};

const post = async function (path, data = null) {
    const uri = getRequestUri(path);

    let response = await axios.post(uri, data);
    console.log("POST: resultStatus: ", response.data.status);

    if (response.data.status !== true) {
        console.log("failed response: ", response);
    }

    return response.data;
};

function getRequestUri(path) {
    return '/api/v1/' + path;
}

export default {
    get,
    post,
};
