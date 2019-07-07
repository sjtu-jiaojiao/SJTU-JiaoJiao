import React, {Component} from 'react';
import { Text, View } from 'react-native';
import { Button, Divider } from "react-native-elements";
import Icon from 'react-native-vector-icons/AntDesign';

export default class HomeScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>交大交交</Text>),
    };

    render() {
        return (
            <View>
                <Button
                    icon={
                        <Icon
                            name="search1"
                            size={22}
                            color="black"
                        />}
                    title="  点我搜索"
                    titleStyle={{color: 'black', fontSize: 17}}
                    buttonStyle={{width: 100, backgroundColor: 'white'}}
                    Type='clear'
                    onPress={() => this.props.navigation.navigate('Search')}
                />
            </View>
        );
    }
}