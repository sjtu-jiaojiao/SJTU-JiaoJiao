import React, {Component} from 'react';
import {
    View, Text, StyleSheet, ScrollView, Alert,
    Image, TouchableOpacity, NativeModules, Dimensions
} from 'react-native';
import Video from 'react-native-video';
import Config from "../../Config";

let ImagePicker = NativeModules.ImageCropPicker;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center'
    },
    button: {
        backgroundColor: 'blue',
        marginBottom: 10
    },
    text: {
        color: 'white',
        fontSize: 20,
        textAlign: 'center'
    }
});


let common_url = 'http://202.120.40.8:30711/v1'; //服务器地址
let dev = "http://202.120.40.8:30711/v1";
class HTTP {
    static uploadImage(url, params) {
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
                    console.warn(responseData);
                    resolve(responseData);
                })
                .catch((err) => {
                    console.warn(err);
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
}

export default class TestPage extends Component {

    constructor() {
        super();
        this.state = {
            image: null,
            images: null
        };
    }

    pickSingleBase64(cropit) {
        ImagePicker.openPicker({
            width: 300,
            height: 300,
            cropping: cropit,
            includeBase64: true,
            includeExif: true,
        }).then(image => {
            console.log('received base64 images');
            this.setState({
                image: {uri: `data:${image.mime};base64,`+ image.data, width: image.width, height: image.height},
                images: null
            });
        }).catch(e => alert(e));
    }

    pickSingle(cropit, circular=false, mediaType) {
        ImagePicker.openPicker({
            width: 500,
            height: 500,
            cropping: cropit,
            cropperCircleOverlay: circular,
            compressImageMaxWidth: 1000,
            compressImageMaxHeight: 1000,
            compressImageQuality: 1,
            compressVideoPreset: 'MediumQuality',
            includeExif: true,
            includeBase64: true,
        }).then(image => {
            console.warn('received base64 image');
            let params = {
                userID: Config.userInfo.userID,
                path: image.path,
            };
            HTTP.uploadImage('/avatar', params)
                .then((response) => {
                    console.warn('成功!');
                    console.warn(response);
                }).catch((err) => {
                    console.warn('失败!');
                    console.warn(err);
                });

            let obj = { };
            HTTP.get(('/file/5d3e68c0d5ef47cec3eb5363'), obj)
                .then((response) => {
                    //alert(response.user[0].userName);
                    this.setState({
                        DataList: response.sellInfo,
                        loaded: true,
                    }, () => {  });
                })
                .catch((error) => console.error(error));
            /*console.warn((Config.fetchPrefix + 'avatar'));
            fetch((Config.fetchPrefix + 'avatar'), {
                method: 'POST',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/x-www-form-urlencoded',
                    Authorization: ('Bearer ' + Config.JaccountToken.token),
                },
                body: ('userID=' + Config.userInfo.userID + '&file=' + image.data)
            })
                .then((response) => {
                    console.warn(response);
                })
                .catch((error) => {
                    console.error(error);
                });
            this.setState({
                image: {uri: `data:${image.mime};base64,`+ image.data, width: image.width, height: image.height},
                images: null
            });
        }).catch(e => {
            console.log(e);
            Alert.alert(e.message ? e.message : e);
        });*/
    })};

    renderVideo(video) {
        console.log('rendering video');
        return (<View style={{height: 300, width: 300}}>
            <Video source={{uri: video.uri, type: video.mime}}
                   style={{position: 'absolute',
                       top: 0,
                       left: 0,
                       bottom: 0,
                       right: 0
                   }}
                   rate={1}
                   paused={false}
                   volume={1}
                   muted={false}
                   resizeMode={'cover'}
                   onError={e => console.log(e)}
                   onLoad={load => console.log(load)}
                   repeat={true} />
        </View>);
    }

    renderImage(image) {
        return <Image style={{width: 300, height: 300, resizeMode: 'contain'}} source={image} />
    }

    renderAsset(image) {
        if (image.mime && image.mime.toLowerCase().indexOf('video/') !== -1) {
            return this.renderVideo(image);
        }

        return this.renderImage(image);
    }

    render() {
        return (<View style={styles.container}>
            <ScrollView>
                {this.state.image ? this.renderAsset(this.state.image) : null}
                {this.state.images ? this.state.images.map(i => <View key={i.uri}>{this.renderAsset(i)}</View>) : null}
            </ScrollView>

            <TouchableOpacity onPress={() => this.pickSingle(true)} style={styles.button}>
                <Text style={styles.text}>Select Single</Text>
            </TouchableOpacity>

        </View>);
    }
}
