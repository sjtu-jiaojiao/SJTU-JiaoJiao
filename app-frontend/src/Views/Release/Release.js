import React, { Component } from 'react';
import {Alert, Text, View} from 'react-native';
import { Button } from 'react-native-elements';
import Config from '../../Config';
import {NavigationActions} from "react-navigation";

export default class ReleaseScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>发布交易信息</Text>)
    };

    render() {
        if (Config.userInfo.userId === -1) {
            Alert.alert(
                '未登录',
                '无法在未登录状态下发布信息，是否切换至登录界面？',
                [
                    {
                        text: '取消',
                        onPress: () => {
                            this.props.navigation.navigate('Home');
                        },
                        style: 'cancel',
                    },
                    {
                        text: '确定', onPress: () => {
                            this.props.navigation.navigate('Login');
                        }
                    },
                ],
                {cancelable: false},
            )
        } else {
            return (
                <View>
                    <Text>
                        {Config.userInfo.userId}
                    </Text>
                </View>
            )
        }
        return (<View/>);
    }
}
