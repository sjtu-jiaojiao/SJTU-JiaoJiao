import React, {Component} from 'React';
import { Text, View, Button } from 'react-native';

export default class HomeScreen extends Component {
    static navigationOptions = {
        title: 'Home',
        headerStyle: {
            backgroundColor: '#f4511e',
        },
        headerTintColor: '#fff',
        headerTitleStyle: {
            fontWeight: 'bold',
        },
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