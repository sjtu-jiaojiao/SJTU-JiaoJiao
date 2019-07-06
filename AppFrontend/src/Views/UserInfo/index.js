import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class UserInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>个人信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是用户信息界面
                </Text>
            </View>
        )
    }
}