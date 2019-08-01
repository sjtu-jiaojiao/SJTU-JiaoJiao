import React, { Component } from 'react';
import { Text, View, Alert } from 'react-native';
import { ListItem, Avatar, Icon } from "react-native-elements";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";

export default class UserScreen extends Component {
    constructor(props) {
        super(props);
        //console.warn(Config)
        this.state = {
            userID: Config.userInfo.userID,
            userName: Config.userInfo.userName,
            studentID: Config.userInfo.studentID,
            studentName: Config.userInfo.studentName,
            telephone: Config.userInfo.telephone,
            avatarID: Config.userInfo.avatarID,
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
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
                        leftAvatar={<Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')}/>}
                        title={'请登录'}
                        titleStyle={{ color: 'black', fontSize: 22 }}
                        onPress={() => {this.props.navigation.navigate('Login')}}
                    />
                    <View style={{height: 35, backgroundColor: '#EFEFF5'}} />
                    <ListItem
                        topDivider
                        bottomDivider
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
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
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
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
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
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
            );
            return null;
        }
        else {
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
                    <ListItem
                        topDivider
                        bottomDivider
                        leftAvatar={() => {
                            if(Config.userInfo.avatarID === 'to_be_changed') {
                                return (
                                    <Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')}/>
                                )
                            } else {
                                return (
                                    <Avatar rounded size='large' source={{uri: (Config.fetchPrefix + 'file/' + this.state.avatarID)}}/>
                                )
                            }
                        }}
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
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
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
                        title={'我的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('MySellInfo')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
                        title={'我的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => this.props.navigation.navigate('MyBuyInfo')}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
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
