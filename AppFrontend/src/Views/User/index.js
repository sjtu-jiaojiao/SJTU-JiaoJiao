import React, { Component } from 'React';
import { Text, View } from 'react-native';

export default class UserScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 25, textAlign: 'center'}}>我</Text>),
    };

    render() {
        return (
            <View>
                <Text>
                    用户界面
                </Text>
            </View>
        )
    }
}