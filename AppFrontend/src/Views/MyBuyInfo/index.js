import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class MyBuyInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>我的求购信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是求购信息界面
                </Text>
            </View>
        )
    }
}