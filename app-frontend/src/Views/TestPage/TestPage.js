import React, {Component} from 'react';
import {
    View, Text, StyleSheet, ScrollView, Alert,
    Image, TouchableOpacity, NativeModules, Dimensions
} from 'react-native';
import Video from 'react-native-video';
import Config from "../../Config";
import {Avatar, withBadge, Badge} from "react-native-elements";


const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        height: 150,
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

const {width, height, scale} = Dimensions.get('window');
let ImagePicker = NativeModules.ImageCropPicker;

export default class TestPage extends Component {

    constructor() {
        super();
        this.state = {
            image: null,
            images: null
        };
        this.keyID = 0;
    }

    pickMultiple() {
            ImagePicker.openPicker({
                multiple: true,
                waitAnimationEnd: false,
                includeExif: true,
                forceJpg: true,
            }).then(selectedImages => {
                this.setState(previousState => {
                    return {
                        image: null,
                        images: previousState.images === null ?
                            selectedImages.map(i => {
                                console.log('received image', i);
                                return {uri: i.path, width: i.width, height: i.height, mime: i.mime};
                            }) :
                            previousState.images.concat(selectedImages.map(i => {
                                console.log('received image', i);
                                return {uri: i.path, width: i.width, height: i.height, mime: i.mime};
                            })),
                    }});
            }).catch(e => {
                console.warn('出错啦！');
                console.warn(e);
            });
    }

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
        return (
            <View>
                <Image style={{
                    width: (1.05 / 5.05 * width),
                    height: (1.05 / 5.05 * width),
                    resizeMode: 'contain',
                    borderColor: 'white',
                    borderWidth: 1,
                }} source={image} />
                <Badge
                    status="error"
                    containerStyle={{ position: 'absolute', top: -4, right: -4 }}
                />
            </View>
        )
    }

    renderAsset(image) {
        if (image.mime && image.mime.toLowerCase().indexOf('video/') !== -1) {
            return this.renderVideo(image);
        }
        return this.renderImage(image);
    }

    render() {
        return (
            <View>
                <View style={{flexDirection: 'row'}}>
                    <TouchableOpacity onPress={this.pickMultiple.bind(this)}>
                        <Image style={{width: (1.05 / 5.05 * width), height: (1.05 / 5.05 * width), resizeMode: 'contain',}} source={require('../../assets/images/addPhotoVideo.jpg')}/>
                    </TouchableOpacity>
                    <ScrollView horizontal={true}>
                        {this.state.images ? this.state.images.map(i => <View key={this.keyID++}>{this.renderAsset(i)}</View>) : null}
                    </ScrollView>
                </View>
            </View>
        );
    }
}
