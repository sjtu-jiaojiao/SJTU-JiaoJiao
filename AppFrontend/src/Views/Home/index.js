import React, {Component} from 'react';
import { Text, View, Button } from 'react-native';

export default class HomeScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>交大交交</Text>),
    };

    render() {
        return (
            <View>
                <Text>
                    这是主页
                </Text>
                <Button
                    title={'Go to test page'}
                    onPress={() => this.props.navigation.navigate('Test')}
                />
            </View>
        )
    }
}