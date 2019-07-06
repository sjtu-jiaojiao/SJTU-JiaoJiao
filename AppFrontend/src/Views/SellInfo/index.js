import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class SellInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 25, textAlign: 'center'}}>我的出售信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是出售信息界面
                </Text>
            </View>
        )
    }
}