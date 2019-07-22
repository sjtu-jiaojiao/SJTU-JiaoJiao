import React, { Component } from 'react';
import {Text, View, TextInput, Alert} from 'react-native';
import {ListItem, Avatar, Button, Icon, Overlay} from "react-native-elements";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";

export default class UserInfoScreen extends Component {
    constructor(props) {
        super(props);
        //console.warn(Config)
        this.state = {
            userId: Config.userInfo.userId,
            userName: Config.userInfo.userName,
            studentId: Config.userInfo.studentId,
            studentName: Config.userInfo.studentName,
            telephone: Config.userInfo.telephone,
            isChangeTelephoneVisible: false,
            isChangeUserNameVisible: false,
        };
        this.changeTelephone = '';
        this.changeUserName = '';
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>个人信息</Text>)
    };

    isTelephoneValid() {
        let length = this.changeTelephone.length;
        return length === 11 && /^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1})|(14[0-9]{1})|)+\d{8})$/.test(this.changeTelephone);
    }

    render() {
        return (
            <View>
                <Overlay
                    isVisible={this.state.isChangeTelephoneVisible}
                    windowBackgroundColor="rgba(255, 255, 255, .5)"
                    overlayBackgroundColor='#EFEFF5'
                    width="auto"
                    height="auto"
                    onBackdropPress={() => {
                        this.changeTelephone = '';
                        this.setState({isChangeTelephoneVisible: false});
                    }}
                >
                    <View style={{width: 300,}}>
                        <Text style={{fontSize: 17}}>请输入新的手机号码</Text>
                        <TextInput
                            autoFocus={true}
                            style={{borderWidth: 1, borderColor: 'black', textAlign: 'center'}}
                            onChangeText={text => this.changeTelephone = text}
                        />
                        <View style={{flexDirection: 'row', justifyContent: 'space-around'}}>
                            <Button
                                title='取消'
                                titleStyle={{color: 'red', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => {
                                    this.changeTelephone = '';
                                    this.setState({isChangeTelephoneVisible: false});
                                }}
                            />
                            <Button
                                title='确认修改'
                                titleStyle={{color: '#298BFF', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => {
                                    if(this.changeTelephone === '') {
                                        Alert.alert(
                                            '电话号码为空',
                                            '电话号码不可以为空，请重新输入',
                                            [
                                                {text: '好', onPress: () => {
                                                    this.setState({isChangeTelephoneVisible: true})
                                                }}
                                            ],
                                            {cancelable: false},
                                        )
                                    }
                                    else if(this.isTelephoneValid() === false) {
                                        Alert.alert(
                                            '不正确的电话号码格式',
                                            this.changeTelephone + '不是正确的电话号码格式，请重新输入',
                                            [
                                                {text: '好', onPress: () => {
                                                        this.setState({isChangeTelephoneVisible: true})
                                                }}
                                            ],
                                            {cancelable: false},
                                        )
                                    }
                                    else {
                                        fetch((Config.fetchPrefix + 'user'), {
                                            method: 'PUT',
                                            headers: {
                                                Accept: 'application/json',
                                                'Content-Type': 'application/x-www-form-urlencoded',
                                                Authorization: ('Bearer ' + Config.JaccountToken.token),
                                            },
                                            body: ('userId=3&telephone=' + this.changeTelephone),
                                        })
                                            .then((response) => {
                                                if(response.ok) {
                                                    Alert.alert(
                                                        '修改成功',
                                                        '成功将电话号码修改为：' + this.changeTelephone,
                                                        [
                                                            {text: '好', onPress: () => {
                                                                    this.setState({isChangeTelephoneVisible: false})
                                                                }}
                                                        ],
                                                        {cancelable: false},
                                                    )
                                                    this.setState({
                                                        telephone: this.changeTelephone,
                                                        isChangeTelephoneVisible: false
                                                    });
                                                    Config.userInfo.telephone = this.changeTelephone;
                                                    this.changeTelephone='';
                                                } else {
                                                    this.changeTelephone='';
                                                    Alert.alert(
                                                        '出错啦',
                                                        '网络可能出了问题，请再试一次吧',
                                                        [
                                                            {text: '好', onPress: () => {
                                                                    this.setState({isChangeTelephoneVisible: false})
                                                                }}
                                                        ],
                                                        {cancelable: false},
                                                    )
                                                }
                                            })
                                            .catch((error) => {
                                                console.error(error);
                                            });
                                    }
                                }}
                            />
                        </View>
                    </View>
                </Overlay>
                <Overlay
                    isVisible={this.state.isChangeUserNameVisible}
                    windowBackgroundColor="rgba(255, 255, 255, .5)"
                    overlayBackgroundColor='#EFEFF5'
                    width="auto"
                    height="auto"
                    onBackdropPress={() => {
                        this.changeUserName='';
                        this.setState({ isChangeUserNameVisible: false });
                    }}
                >
                    <View style={{width: 300,}}>
                        <Text style={{fontSize: 17}}>请输入新的用户名</Text>
                        <TextInput
                            autoFocus={true}
                            style={{borderWidth: 1, borderColor: 'black', textAlign: 'center'}}
                            onChangeText={text => this.changeUserName = text}
                        />
                        <View style={{flexDirection: 'row', justifyContent: 'space-around'}}>
                            <Button
                                title='取消'
                                titleStyle={{color: 'red', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => {
                                    this.changeUserName = '';
                                    this.setState({isChangeUserNameVisible: false});
                                }}
                            />
                            <Button
                                title='确认修改'
                                titleStyle={{color: '#298BFF', fontSize: 17, fontWeight: 'bold'}}
                                buttonStyle={{backgroundColor: '#EFEFF5'}}
                                raised={true}
                                onPress={() => {
                                    if(this.changeUserName === '') {
                                        Alert.alert(
                                            '用户名为空',
                                            '用户名不可以为空，请重新输入',
                                            [
                                                {text: '好', onPress: () => {
                                                        this.setState({isChangeUserNameVisible: true})
                                                    }}
                                            ],
                                            {cancelable: false},
                                        )
                                    }
                                    else if(this.changeUserName.length > 32) {
                                        Alert.alert(
                                            '用户名过长',
                                            '用户名长度不可以超过32个字符，请重新输入',
                                            [
                                                {text: '好', onPress: () => {
                                                        this.setState({isChangeUserNameVisible: true})
                                                }}
                                            ],
                                            {cancelable: false},
                                        )
                                    }
                                    else {
                                        fetch((Config.fetchPrefix + 'user'), {
                                            method: 'PUT',
                                            headers: {
                                                Accept: 'application/json',
                                                'Content-Type': 'application/x-www-form-urlencoded',
                                                Authorization: ('Bearer ' + Config.JaccountToken.token),
                                            },
                                            body: ('userId=3&userName=' + this.changeUserName),
                                        })
                                            .then((response) => {
                                                if(response.ok) {
                                                    Alert.alert(
                                                        '修改成功',
                                                        '成功将用户名修改为：' + this.changeUserName,
                                                        [
                                                            {text: '好', onPress: () => {
                                                                    this.setState({isChangeUserNameVisible: false})
                                                                }}
                                                        ],
                                                        {cancelable: false},
                                                    )
                                                    this.setState({
                                                        userName: this.changeUserName,
                                                        isChangeUserNameVisible: false
                                                    });
                                                    Config.userInfo.userName = this.changeUserName;
                                                    this.changeUserName='';
                                                } else {
                                                    this.changeUserName='';
                                                    Alert.alert(
                                                        '出错啦',
                                                        '网络可能出了问题，请再试一次吧',
                                                        [
                                                            {text: '好', onPress: () => {
                                                                    this.setState({isChangeUserNameVisible: false})
                                                                }}
                                                        ],
                                                        {cancelable: false},
                                                    )
                                                }
                                            })
                                            .catch((error) => {
                                                console.error(error);
                                            });
                                    }
                                }}
                            />
                        </View>
                    </View>
                </Overlay>
                <View style={{height: 200}}>
                    <View style={{alignItems: 'center', justifyContent: 'center'}}>
                        <View style={{height: 55}}/>
                        <Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')} />
                        <Text style={{fontSize: 15}}>点击修改头像</Text>
                    </View>
                </View>
                <ListItem
                    topDivider
                    bottomDivider
                    title={'用户名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.userName}</Text>}
                    onPress={() => this.setState({isChangeUserNameVisible: true})}
                />
                <ListItem
                    bottomDivider
                    title={'手机'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.telephone}</Text>}
                    onPress={() => this.setState({isChangeTelephoneVisible: true})}
                />
                <ListItem
                    bottomDivider
                    title={'学号'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'grey', fontSize: 17 }}>{this.state.studentId}</Text>}
                />
                <ListItem
                    bottomDivider
                    title={'真实姓名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'grey', fontSize: 17 }}>{this.state.studentName}</Text>}
                />
                <View style={{height: 15}}/>
                <Button
                    title='退出登录'
                    titleStyle={{color: 'white', fontSize: 17}}
                    buttonStyle={{backgroundColor: 'red'}}
                    containerStyle={{width: 160, marginLeft: 120}}
                    raised={true}
                    onPress={() => Alert.alert(
                        '退出登录',
                        '您确定要退出登录吗？',
                        [
                            {
                                text: '取消',
                                onPress: () => console.log('Cancel Pressed'),
                                style: 'cancel',
                            },
                            {
                                text: '确定', onPress: () => {
                                    Config.JaccountToken={};
                                    Config.userInfo = {
                                        userId: -1,
                                        userName: '',
                                        avatarId: '',
                                        telephone: '',
                                        studentId: '',
                                        studentName: '',
                                    };
                                    this.props.navigation.reset([NavigationActions.navigate({ routeName: 'User' })], 0);
                                }
                            },
                        ],
                        {cancelable: false},
                    )}
                />
            </View>
        )
    }
}
