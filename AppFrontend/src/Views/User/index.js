import React, { Component } from 'React';
import { Text, View, FlatList } from 'react-native';
import { ListItem, Avatar } from "react-native-elements";
import Icon from "react-native-vector-icons/AntDesign";

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
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <View style={{height: 20, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        leftAvatar={<Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />}
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={this.state.userName}
                        titleStyle={{ color: 'black', fontSize: 25 }}
                        subtitle={this.state.email}
                        subtitleStyle={{ color: 'black', fontSize: 15 }}
                        onPress={() => this.props.navigation.navigate('UserInfo')}
                    />
                    <View style={{height: 35, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('MySellInfo')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('MyBuyInfo')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的历史交易记录'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('MyHistoryInfo')}
                    />
                    <ListItem
                        containerStyle={{height: 1000, backgroundColor: '#EFEFF5'}}
                    />
                </View>
            )
        }
        else {
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <View style={{height: 20, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        leftAvatar={<Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />}
                        title={'请登录'}
                        titleStyle={{ color: 'black', fontSize: 22 }}
                        onPress={() => this.props.navigation.navigate('Login')}
                    />
                    <View style={{height: 35, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的历史交易记录'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => alert('请先登录！')}
                    />
                    <ListItem
                        containerStyle={{height: 1000, backgroundColor: '#EFEFF5'}}
                    />
                </View>
            )
        }
    }
}