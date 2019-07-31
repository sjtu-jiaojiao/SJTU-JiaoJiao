import React, {Component} from 'react';
import {
    StyleSheet,
    View,
    Text,
    ScrollView,
    RefreshControl, Image, TouchableOpacity,
} from 'react-native';
import jwt_decode from 'jwt-decode';
import Config from "../../Config";

export default class RefreshControlDemo extends Component {
    static navigationOptions = {
        title: 'RefreshControl',

    };

    state = {
        loaded:0,
        isRefreshing: false,
        data: Array.from(new Array(20)).map((val, i) => ({text: '初始化： ' + i, clicks: 0})),
    };

    _onRefresh = () => {
        this.setState({isRefreshing: true});
        setTimeout(() => {
            // prepend 10 items
            const rowData = Array.from(new Array(10))
                .map((val, i) => ({
                    text: '第几次加载： ' + (+this.state.loaded + i),
                    clicks: 0,
                }))
                .concat(this.state.data);

            this.setState({
                loaded: this.state.loaded + 10,
                isRefreshing: false,
                data: rowData,
            });
        }, 3000);
    };

    render() {
        let imageUri = 'http://202.120.40.8:30711/v1/file/5d41bc5bc286640d0c7ec63e';
        let imageWidth, imageHeight;
        Image.getSize(imageUri)
            .then((width, height) => {
                imageWidth = 0.85 * width;
                imageHeight = height * imageWidth / width;
            })
            .catch((error) => {
                console.warn(error);
            });
        return (
            <TouchableOpacity onLongPress={() => {}}>
                <Image style={{
                    width: imageWidth,
                    height: imageHeight,
                    resizeMode: 'contain',
                    borderColor: 'white',
                    borderWidth: 1,
                }} source={{uri: imageUri}} />
            </TouchableOpacity>
        );
        /*return (
            <ScrollView
                style={{flex:1}}
                refreshControl={
                    <RefreshControl
                        refreshing={this.state.isRefreshing}
                        onRefresh={this._onRefresh}
                        tintColor="#ff0000"
                        title="Loading..."
                        titleColor="#00ff00"
                        size={0}
                        progressViewOffset={30}
                        colors={['#0000ff','#ff0000', '#00ff00', ]}
                        progressBackgroundColor="#ffff00"
                    />
                }>
                    <TouchableOpacity onLongPress={() => {}}>
                        <Image source={{uri: 'http://202.120.40.8:30711/v1/file/5d41bc5bc286640d0c7ec63e'}} />
                    </TouchableOpacity>
                 <View>
                    {
                        this.state.data.map((row, ii) => {
                            return (<Text>{row.text}</Text>);
                        })
                    }
                </View>
            </ScrollView>
        );*/
    }
}
0
