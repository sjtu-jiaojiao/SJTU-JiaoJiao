import React, { Component } from 'React';
import { Text, View } from 'react-native';

export default class ReleaseScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 25, textAlign: 'center'}}>发布交易信息</Text>)
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