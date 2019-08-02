import React, {Component} from 'react';
import Config from "../Config";

let common_url = 'http://202.120.40.8:30711/v1'; //服务器地址

let dev = "http://202.120.40.8:30711/v1";

export default class HTTP {
    static addAvatar(url, params) {
        return new Promise(function (resolve, reject) {
            let formData = new FormData();
            for (let key in params) {
                formData.append(key, params[key]);
            }
            let file = {uri: params.path, type: 'application/octet-stream', name: 'image.jpg'};
            formData.append("file", file);
            fetch(common_url + url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'multipart/form-data;charset=utf-8',
                    Authorization: ('Bearer ' + Config.JaccountToken.token),
                },
                body: formData,
            }).then((response) => response.json())
                .then((responseData) => {
                    //console.warn(responseData);
                    resolve(responseData);
                })
                .catch((err) => {
                    //console.warn(err);
                    reject(err);
                });
        });
    }

    static addContent(url, params) {
        return new Promise(function (resolve, reject) {
            let formData = new FormData();
            for (let key in params) {
                formData.append(key, params[key]);
            }
            let content = {uri: params.path, type: 'application/octet-stream', name: 'image.jpg'};
            formData.append("content", content);
            fetch(common_url + url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'multipart/form-data;charset=utf-8',
                    Authorization: ('Bearer ' + Config.JaccountToken.token),
                },
                body: formData,
            }).then((response) => response.json())
                .then((responseData) => {
                    //console.warn(responseData);
                    resolve(responseData);
                })
                .catch((err) => {
                    //console.warn(err);
                    reject(err);
                });
        });
    }

    static get(url, params) {
        // 将后台接口的公共部分拼接进去
        url = dev + url;
        //判断有木有参数
        if (params) {
            // 定一个空数组
            let paramsArray = [];
            //  拆分对象
            Object.keys(params).forEach(key =>
                paramsArray.push(key + "=" + params[key])
            );
            // 判断是否地址拼接的有没有 ？,当没有的时候，使用 ？拼接第一个参数，如果有参数拼接，则用&符号拼接后边的参数
            if (url.search(/\?/) === -1) {
                url = url + "?" + paramsArray.join("&");
            } else {
                url = url + "&" + paramsArray.join("&");
            }
        }
        // 返回一个promise
        return new Promise((resolve, reject) => {
            //fetch请求
            fetch(url, { method: "GET" })
                .then(response => response.json())
                .then(resulet => {
                    resolve(resulet);
                })
                .catch(error => {
                    reject(error);
                });
        });
    }

    static addInfo(url, formData) {
        //console.warn((common_url + url));
        //console.warn(formData);
        return new Promise(function (resolve, reject) {
            fetch((common_url + url), {
                method: 'POST',
                headers: {
                    'Content-Type': 'multipart/form-data;charset=utf-8',
                    Authorization: ('Bearer ' + Config.JaccountToken.token),
                },
                body: formData,
            }).then((response) => response.json())
                .then((responseData) => {
                    //console.warn(responseData);
                    resolve(responseData);
                })
                .catch((err) => {
                    //console.warn(err);
                    reject(err);
                });
        });
    }

}
