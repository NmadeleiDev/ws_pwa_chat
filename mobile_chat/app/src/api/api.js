const httpModule = require("tns-core-modules/http");

const get = async function (path) {
    const uri = buildRequestUri(path);
    let result;
    let response;

    try {
        response = await httpModule.getJSON(uri);
    } catch (e) {
        console.log("get request error: ", err)
        return {
            error: 'network error',
            data: null,
        }
    }
    response = response.content.toJSON();

    if (response.status === false) {
        result = {
            error: response.data,
            data: null,
        }
    } else {
        result = {
            error: null,
            data: response.data,
        }
    }

    return result;
};

const post = async function (path, data = null, eventDate = null) {
    const uri = buildRequestUri(path);
    let response;
    let result;
    let headers;

    console.log("Requesting ", uri)
    console.log("Data ", data)
    if (eventDate !== null) {
        headers = { "Content-Type": "application/json", "Event-date": eventDate, "mobile": 'true' }
    } else {
        headers = { "Content-Type": "application/json" }
    }

    try {
        response = await httpModule.request({
            url: uri,
            method: "POST",
            headers: headers,
            content: JSON.stringify(data)
          });
    } catch (e) {
        console.log("post request error: ", e)
        return {
            error: 'network error',
            data: null,
        }
    }
    response = response.content.toJSON();

    if (response.status === false) {
        result = {
            error: response.data,
            data: null,
        }
    } else {
        result = {
            error: null,
            data: response.data,
        }
    }
    console.log("Result: ", result);
    return result;
};

function buildRequestUri(path) {
    return 'http://192.168.1.83:8080/api/v1/' + path;
}

export default {
    get,
    post,
};
