import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class ReleaseScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>发布交易信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    发布交易信息
                </Text>
            </View>
        )
    }
}