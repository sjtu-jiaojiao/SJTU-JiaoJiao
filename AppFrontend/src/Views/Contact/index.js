import React, { Component } from 'React';
import { Text, View } from 'react-native';

export default class ContactScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 25, textAlign: 'center'}}>消息</Text>)
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