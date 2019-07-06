import React, { Component } from 'React';
import { Text, View, FlatList } from 'react-native';
import { ListItem, Avatar } from "react-native-elements";

export default class UserScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isLogin: false,
            userName: 'chiangel',
            realName: '林江浩',
            gender: '男',
            email: 'chiangel.ljh@gmail.com',
            phoneNumber: '15221278083',
        }
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>我</Text>),
    };

    render() {
        if (this.state.isLogin) {
            return (
                <View>
                    <View style={{height: 20}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        chevron
                        leftAvatar={<Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />}
                        title={this.state.userName}
                        titleStyle={{ color: 'black', fontSize: 25 }}
                        subtitle={this.state.email}
                        subtitleStyle={{ color: 'black', fontSize: 15 }}
                    />
                </View>
            )
        }
        else {
            return (
                <View>
                    <View style={{height: 20}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        chevron
                        leftAvatar={<Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />}
                        title={this.state.userName}
                        titleStyle={{ color: 'black', fontSize: 25 }}
                        subtitle={this.state.email}
                        subtitleStyle={{ color: 'black', fontSize: 15 }}
                    />
                </View>
            )
        }
    }
}