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
                        onPress={() => this.props.navigation.navigate('UserInfo')}
                    />
                    <View style={{height: 35}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        chevron
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('SellInfo')}
                    />
                    <ListItem
                        bottomDivider
                        chevron
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('SellInfo')}
                    />
                    <ListItem
                        bottomDivider
                        chevron
                        title={'我的历史交易记录'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('HistoryInfo')}
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
                        title={'请登录'}
                        titleStyle={{ color: 'black', fontSize: 22 }}
                        onPress={() => {
                            alert('即将通过Jaccount登录！');
                            this.setState(previousState => {
                                return {isLogin: !previousState.isLogin};
                            })}}
                    />
                    <View style={{height: 35}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        chevron
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                    <ListItem
                        bottomDivider
                        chevron
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                    <ListItem
                        bottomDivider
                        chevron
                        title={'我的历史交易记录'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                </View>
            )
        }
    }
}