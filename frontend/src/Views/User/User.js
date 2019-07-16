import React, { Component } from 'React';
import { Text, View, Alert } from 'react-native';
import { ListItem, Avatar } from "react-native-elements";
import Icon from "react-native-vector-icons/AntDesign";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";

export default class UserScreen extends Component {
    constructor(props) {
        super(props);
        //console.warn(Config)
        this.state = {
            userId: Config.userInfo.userId,
            userName: Config.userInfo.userName,
            studentId: Config.userInfo.studentId,
            studentName: Config.userInfo.studentName,
            telephone: Config.userInfo.telephone,
        };
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>我</Text>),
    };

    render() {
        //console.warn(Config.JaccountToken);
        if (!Config.JaccountToken.hasOwnProperty('status')) {
            //console.warn(Config.JaccountToken);
            //console.warn(Config.JaccountToken);
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
                        onPress={() => {this.props.navigation.navigate('Login')}}
                    />
                    <View style={{height: 35, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => Alert.alert(
                            '未登录',
                            '请先点击登录哦~',
                            [
                                {text: '好'},
                            ],
                            {cancelable: false},
                        )}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => Alert.alert(
                            '未登录',
                            '请先点击登录哦~',
                            [
                                {text: '好'},
                            ],
                            {cancelable: false},
                        )}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Icon name='rightcircleo' size={20} color={'grey'} />}
                        title={'我的历史交易记录'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => Alert.alert(
                            '未登录',
                            '请先点击登录哦~',
                            [
                                {text: '好'},
                            ],
                            {cancelable: false},
                        )}
                    />
                    <ListItem
                        containerStyle={{height: 1000, backgroundColor: '#EFEFF5'}}
                    />
                </View>
            )
        }
        else if (Config.JaccountToken.status === 2) {
            Alert.alert(
                '登录失败',
                '似乎没有登录成功，请再试一次吧~',
                [
                    {text: '好', onPress: () => {
                            Config.JaccountToken={};
                            this.props.navigation.reset([NavigationActions.navigate({ routeName: 'User' })], 0);
                        }}
                ],
                {cancelable: false},
            )
            return null;
        }
        else {
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
                        subtitle={this.state.telephone}
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
    }
}
