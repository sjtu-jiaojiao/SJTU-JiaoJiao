import React, { Component } from 'react';
import { Text, View } from "react-native";

export default class ContactScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>消息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是信息界面
                </Text>
            </View>
        )
    }
}