import axios from 'axios'

const get = async function (path) {
    const uri = getRequestUri(path);
    let resultStatus;

    let response = await axios.get(uri);
    resultStatus = response.data.status;

    if (resultStatus !== true) {
        console.log(response);
    }

    return response.data;
};

const post = async function (path, data = null) {
    const uri = getRequestUri(path);

    let response = await axios.post(uri, {data: data});

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
