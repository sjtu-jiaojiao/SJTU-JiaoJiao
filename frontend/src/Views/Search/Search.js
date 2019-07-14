import React, { Component } from 'react';
import {Text, View, FlatList, Alert} from 'react-native';
import {SearchBar, Button, ListItem} from "react-native-elements";
import Config from '../../Config';
import {NavigationActions} from "react-navigation";

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
            searchType: 'User',
            showSearchItem: false,
            UserList: null,
        }
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}> </Text>)
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => (
        <ListItem
            bottomDivider
            title={item.userName}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    );

    updateSearch = (searchText) => {
        this.setState({ searchText: searchText });
    };

    cancelSearch = (searchText) => {
        this.setState({
            showSearchItem: false,
            searchText: '',
        });
    };

    startSearch = () => {
        if(this.state.searchText !== '') {
            let obj = { userName: this.state.searchText };
            Http.get('/user', obj)
                .then((response) => {
                    this.setState({
                        UserList: response.user,
                    });
                    //console.warn(this.state.UserList);
                })
                .catch((error) => console.error(error));
        }
    };

    render() {
        const { searchText } = this.state.searchText;
        if (this.state.UserList !== null) {
            return (
                <View>
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            containerStyle={{flex: 4}}
                            placeholder={'搜索'}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'white', fontSize: 17}}
                            buttonStyle={{backgroundColor: 'steelblue'}}
                            containerStyle={{flex: 1}}
                            raised={true}
                            onPress={() => this.startSearch()}
                        />
                    </View>
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.UserList}
                        renderItem={this.renderItem}
                    />
                </View>
            )
        }
        else {
            return (
                <View>
                    <View style={{flexDirection: 'row'}}>
                        <SearchBar
                            containerStyle={{flex: 4}}
                            placeholder={'搜索'}
                            onChangeText={ this.updateSearch }
                            onClear={this.cancelSearch}
                            value={ this.state.searchText }
                        />
                        <Button
                            title='搜  索'
                            titleStyle={{color: 'white', fontSize: 17}}
                            buttonStyle={{flex: 1}}
                            containerStyle={{backgroundColor: 'steelblue', flex: 1, alignItems: 'center'}}
                            raised={true}
                            onPress={() => this.startSearch()}
                        />
                    </View>
                </View>
            )
        }
    }
}
