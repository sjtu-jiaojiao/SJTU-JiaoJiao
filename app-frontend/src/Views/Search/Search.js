import React, { Component } from 'react';
import {Text, View, FlatList, Alert, TextInput, StyleSheet, TouchableOpacity} from 'react-native';
import {SearchBar, Button, ListItem, ButtonGroup, Avatar} from "react-native-elements";
import Config from '../../Config';
import {isUserNameValid} from '../../Utils/CheckValidity';

let dev = "http://202.120.40.8:30711/v1";

class Http {
    // 静态方法
    static get(url, params) {
        // 将后台接口的公共部分拼接进去
        url = dev + url;
        //判断有木有参数
        if (params) {
            // 定一个空数组
            let paramsArray = [];
            //  拆分对象
            Object.keys(params).forEach(key =>
                paramsArray.push(key + "=" + params[key])
            );
            // 判断是否地址拼接的有没有 ？,当没有的时候，使用 ？拼接第一个参数，如果有参数拼接，则用&符号拼接后边的参数
            if (url.search(/\?/) === -1) {
                url = url + "?" + paramsArray.join("&");
            } else {
                url = url + "&" + paramsArray.join("&");
            }
        }
        // 返回一个promise
        return new Promise((resolve, reject) => {
            //fetch请求
            fetch(url, { method: "GET" })
                .then(response => response.json())
                .then(resulet => {
                    resolve(resulet);
                })
                .catch(error => {
                    reject(error);
                });
        });
    }
}

export default class SearchScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            searchText: '',
            DataList: null,
            selectedIndex: 1,
            loaded: false,
        };
        this.updateIndex = this.updateIndex.bind(this);
    };

    updateIndex (selectedIndex) {
        this.setState({
            selectedIndex: selectedIndex,
            searchText: '',
            DataList: [],
            loaded: false,
        })
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}> </Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => {
        let renderContent, routeDest, params;
        let userName, userID, sellInfoID, buyInfoID, header, infoType;
        switch (this.state.selectedIndex) {
            case 0:
                renderContent = item.userName;
                routeDest = 'UserInfoForOthers';
                userName = item.userName;
                userID = item.userID;
                params = { userName, userID };
                break;
            case 1:
                renderContent = item.goodName;
                routeDest = 'GoodInfo';
                infoType = 'sellInfo';
                sellInfoID = item.sellInfoID;
                header = ('出售：' + item.goodName);
                params = { infoType, sellInfoID, header };
                break;
            case 2:
                renderContent = item.goodName;
                routeDest = 'GoodInfo';
                infoType = 'buyInfo';
                buyInfoID = item.buyInfoID;
                header = ('求购：' + item.goodName);
                params = { infoType, buyInfoID, header };
                break;
        }
        return (
            <TouchableOpacity onPress={() => this.props.navigation.push(routeDest, params)}>
                <ListItem
                    bottomDivider
                    containerStyle={{height: 50}}
                    rightIcon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/right.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    title={
                        <Text numberOfLines={1} style={{color: 'black', fontSize: 20}}>{renderContent}</Text>
                    }
                />
            </TouchableOpacity>
        );
    };

    updateSearch = (searchText) => {
        if (this.state.loaded === true && this.state.DataList === undefined) {
            this.setState({
                searchText: searchText,
                loaded: false,
                DataList: null,
            });
        } else {
            this.setState({searchText: searchText});
        }
    };

    cancelSearch = (searchText) => {
        this.setState({
            searchText: '',
            DataList: [],
            loaded: false,
        });
    };

    startSearch = () => {
        //console.warn('Search ' + this.state.searchText + ' in ' + this.state.selectedIndex);
        if (this.state.selectedIndex === 0) {
            let obj = { userName: this.state.searchText };
            //console.warn(obj);
            Http.get('/user', obj)
                .then((response) => {
                    //alert(response.user[0].userName);
                    this.setState({
                        DataList: response.user,
                        loaded: true,
                    }, () => {  });
                })
                .catch((error) => console.error(error));
        } else if (this.state.selectedIndex === 1) {
            let obj = { goodName: this.state.searchText };
            //console.warn(obj);
            Http.get('/sellInfo', obj)
                .then((response) => {
                    //alert(response.user[0].userName);
                    this.setState({
                        DataList: response.sellInfo,
                        loaded: true,
                    }, () => {  });
                })
                .catch((error) => console.error(error));
        } else if (this.state.selectedIndex === 2) {
            let obj = { goodName: this.state.searchText };
            //console.warn(obj);
            Http.get('/buyInfo', obj)
                .then((response) => {
                    //alert(response.user[0].userName);
                    this.setState({
                        DataList: response.buyInfo,
                        loaded: true,
                    }, () => {  });
                })
                .catch((error) => console.error(error));
        }
    };

    render() {
        const buttons = ['用户', '出售信息', '求购信息'];
        const { selectedIndex } = this.state;
        const { searchText } = this.state;
        if (this.state.loaded === false) {
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <ButtonGroup
                        onPress={this.updateIndex}
                        selectedIndex={selectedIndex}
                        buttons={buttons}
                        containerStyle={{height: 35}}
                        disabledTextStyle={{fontSize: 17}}
                        selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                    />
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            searchIcon={false}
                            lightTheme={true}
                            containerStyle={{
                                flex: 4,
                                height: 55,
                                backgroundColor: 'white',
                                borderWidth: 0, //no effect
                                shadowColor: 'white', //no effect
                                borderBottomColor: 'transparent',
                                borderTopColor: 'transparent',
                            }}
                            inputContainerStyle={{
                                marginTop: -9,
                                height: 55,
                                backgroundColor: 'white',
                            }}
                            placeholder={'搜索' + buttons[this.state.selectedIndex]}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            clearIcon={
                                <Avatar
                                    rounded size='small'
                                    source={require('../../assets/icons/cancel.png')}
                                />
                            }
                            cancelIcon={false}
                            onPress={() => {
                                if(this.state.searchText === '') {
                                    Alert.alert(
                                        '输入为空',
                                        '搜索输入不可以为空，请重新输入',
                                        [
                                            {text: '好', onPress: () => {}}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else if(this.state.selectedIndex === 0 && !isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且只能包含汉字、字母、数字和特殊符号“_”、“.”',
                                        [
                                            {text: '好', onPress: () => {
                                                this.setState({
                                                    DataList: [],
                                                    loaded: false,
                                                });
                                            }}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else {
                                    this.startSearch();
                                }
                            }}
                        />
                    </View>
                    <View style={{height: 10}} />
                </View>
            )
        }
        else if (this.state.DataList === undefined) {
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <ButtonGroup
                        onPress={this.updateIndex}
                        selectedIndex={selectedIndex}
                        buttons={buttons}
                        containerStyle={{height: 35}}
                        disabledTextStyle={{fontSize: 17}}
                        selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                    />
                    <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                        <SearchBar
                            searchIcon={false}
                            lightTheme={true}
                            containerStyle={{
                                flex: 4,
                                height: 55,
                                backgroundColor: 'white',
                                borderWidth: 0, //no effect
                                shadowColor: 'white', //no effect
                                borderBottomColor: 'transparent',
                                borderTopColor: 'transparent',
                            }}
                            inputContainerStyle={{
                                marginTop: -9,
                                height: 55,
                                backgroundColor: 'white',
                            }}
                            placeholder={'搜索' + buttons[this.state.selectedIndex]}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            clearIcon={
                                <Avatar
                                    rounded size='small'
                                    source={require('../../assets/icons/cancel.png')}
                                />
                            }
                            cancelIcon={false}
                            onPress={() => {
                                if(this.state.searchText === '') {
                                    Alert.alert(
                                        '输入为空',
                                        '搜索输入不可以为空，请重新输入',
                                        [
                                            {text: '好', onPress: () => {}}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else if(this.state.selectedIndex === 0 && !isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且只能包含汉字、字母、数字和特殊符号“_”、“.”',
                                        [
                                            {text: '好', onPress: () => {
                                                this.setState({
                                                    DataList: [],
                                                    loaded: false,
                                                });
                                            }}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else {
                                    this.startSearch();
                                }
                            }}
                        />
                    </View>
                    <View style={{height: 10}} />
                    <View style={{alignItems: 'center', justifyContent: 'center'}}>
                        <Text style={{fontSize: 17}}>没有匹配的{buttons[this.state.selectedIndex]}，换个关键词试试？</Text>
                    </View>
                    <View style={{height: 10}} />
                </View>
            )
        }
        else {
            //console.warn('Have DataList Now!');
            //alert(this.state.DataList[0].userName);
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <ButtonGroup
                        onPress={this.updateIndex}
                        selectedIndex={selectedIndex}
                        buttons={buttons}
                        containerStyle={{height: 35}}
                        disabledTextStyle={{fontSize: 17}}
                        selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                    />
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            searchIcon={false}
                            lightTheme={true}
                            containerStyle={{
                                flex: 4,
                                height: 55,
                                backgroundColor: 'white',
                                borderWidth: 0, //no effect
                                shadowColor: 'white', //no effect
                                borderBottomColor: 'transparent',
                                borderTopColor: 'transparent',
                            }}
                            inputContainerStyle={{
                                marginTop: -9,
                                height: 55,
                                backgroundColor: 'white',
                            }}
                            placeholder={'搜索' + buttons[this.state.selectedIndex]}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'steelblue', fontSize: 17, fontWeight: 'bold'}}
                            containerStyle={{flex: 1}}
                            buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                            raised={true}
                            clearIcon={
                                <Avatar
                                    rounded size='small'
                                    source={require('../../assets/icons/cancel.png')}
                                />
                            }
                            cancelIcon={false}
                            onPress={() => {
                                if(this.state.searchText === '') {
                                    Alert.alert(
                                        '输入为空',
                                        '搜索输入不可以为空，请重新输入',
                                        [
                                            {text: '好', onPress: () => {}}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else if(this.state.selectedIndex === 0 && !isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且只能包含汉字、字母、数字和特殊符号“_”、“.”',
                                        [
                                            {text: '好', onPress: () => {
                                                this.setState({
                                                    DataList: [],
                                                    loaded: false,
                                                });
                                            }}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                                else {
                                    this.startSearch();
                                }
                            }}
                        />
                    </View>
                    <View style={{height: 10}} />
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.DataList}
                        renderItem={this.renderItem}
                    />
                </View>
            )
        }
    }
}
