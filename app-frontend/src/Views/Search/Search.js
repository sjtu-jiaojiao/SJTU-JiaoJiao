import React, { Component } from 'react';
import {Text, View, FlatList, Alert, TextInput, StyleSheet} from 'react-native';
import {SearchBar, Button, ListItem, ButtonGroup, Avatar} from "react-native-elements";
import Config from '../../Config';

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

    parseTitle(item) {
        switch (this.state.selectedIndex) {
            case 0:
                return item.userName;
            case 1:
                return item.goodName;
            case 2:
                return item.goodName;
        }
    }

    renderItem = ({ item }) => {
        //alert(item.userName);
        return (
            <ListItem
                bottomDivider
                containerStyle={{height: 50}}
                title={
                    <Text numberOfLines={1} style={{color: 'black', fontSize: 20}}>{this.parseTitle(item)}</Text>
                }
            />
        );
    };

    updateSearch = (searchText) => {
        this.setState({ searchText: searchText });
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
            console.warn('Search SellInfo');
        } else {
            console.warn('Search BuyInfo')
        }
        // 注意，如果返回的list为空，要提示用户没有他想要的！
    };

    isUserNameValid() {
        return /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){0,31}$/.test(this.state.searchText);
    }

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
                                else if(this.state.selectedIndex === 0 && !this.isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且以字母开头、可带数字、“_”、“.”',
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
                                else if(this.state.selectedIndex === 0 && !this.isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且以字母开头、可带数字、“_”、“.”',
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
                        <Text style={{fontSize: 17}}>没有你要的用户，换个关键词试试？</Text>
                    </View>
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
                                else if(this.state.selectedIndex === 0 && !this.isUserNameValid(this.state.searchText)) {
                                    Alert.alert(
                                        '用户名不合法',
                                        '要求用户名不超过32位，且以字母开头、可带数字、“_”、“.”',
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
