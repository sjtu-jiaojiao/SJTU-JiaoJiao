import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class SellInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>出售信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    浏览出售信息
                </Text>
            </View>
        )
    }
}