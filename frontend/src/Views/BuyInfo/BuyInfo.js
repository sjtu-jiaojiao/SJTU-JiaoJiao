import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class BuyInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>求购信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    浏览求购信息
                </Text>
            </View>
        )
    }
}