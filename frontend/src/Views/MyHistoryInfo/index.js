import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class MyHistoryInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>我的历史交易</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是历史交易信息
                </Text>
            </View>
        )
    }
}