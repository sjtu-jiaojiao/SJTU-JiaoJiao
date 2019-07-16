import React, { Component } from 'react';
import {Text, View, FlatList, Alert} from 'react-native';
import {ListItem, Avatar, Button} from "react-native-elements";
import Icon from "react-native-vector-icons/AntDesign";
import Config from "../../Config";
import {NavigationActions} from "react-navigation";
import axios from "axios";

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
        };
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}>个人信息</Text>)
    };

    updateUserInfo() {
        let formData = new FormData();
        formData.append('userId', 1);
        formData.append('telephone', '15221278083');
        formData.append('status', 1);
        /*
        let postxml = new XMLHttpRequest();
        postxml.onreadystatechange = (e) => {
            if (postxml.readyState !== 4) {
                return;
            }
            if (postxml.status === 200) {
                alert("请求成功！");
                //this.setState({responseText: postxml.responseText})
                console.warn(postxml.responseText);

            } else {
                alert("请求失败！");
            }
        };
        postxml.open("POST","ajax_test.asp",true);
        postxml.setRequestHeader("Content-type","application/json");
        postxml.send("userId=1&telephone=15221278083&status=1");*/

        axios.post('http://202.120.40.8:30711/v1/user',
            {
                'userId': '1',
                'telephone': '15221278083',
                'status': '1',
            },
            {
                headers: {
                    'Authorization': 'Bearer ' + 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMxMzY3MTQsImlkIjo0LCJyb2xlIjoxfQ.9UCj-BC-hL1KVRBs8KCPVE4LSKDriicCuCZ_s495_Pg',
                    'Content-Type': 'application/x-www-form-urlencoded',
                }
            }
        )
            .then(res => console.warn(res))
            .catch(e => console.warn(e));

        /*axios.post(
            '202.120.40.8:30711/v1/user',
            {
                'userId': '1',
                'telephone': '15221278083',
                'status': '1',
            }
        )
            .then(function (response) {
                // handle success
                console.warn(response);
            })
            .catch(function (error) {
                // handle error
                console.warn(error);
            })
            .finally(function () {
                // always executed
            });*/
        /*fetch('http://202.120.40.8:30711/v1/user', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/x-www-form-urlencoded',
                Authorization: ('Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMxMzY3MTQsImlkIjo0LCJyb2xlIjoxfQ.9UCj-BC-hL1KVRBs8KCPVE4LSKDriicCuCZ_s495_Pg'),
            },
            body: 'userId=1&telephone=15221278083&status=1',
        })
            .then((response) => {
                console.warn(response);
            })
            .catch((error) => {
                console.error(error);
            });*/
    }

    render() {
        return (
            <View>
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
                    chevron
                    title={'用户名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.userName}</Text>}
                    onPress={ () => {
                        this.updateUserInfo();
                    }}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'手机'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.telephone}</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'学号'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.studentId}</Text>}
                />
                <ListItem
                    bottomDivider
                    chevron
                    title={'真实姓名'}
                    titleStyle={{ color: 'black', fontSize: 17 }}
                    rightElement={<Text style={{ color: 'steelblue', fontSize: 17 }}>{this.state.studentName}</Text>}
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
