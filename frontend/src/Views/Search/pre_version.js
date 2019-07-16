import React, { Component } from 'react';
import { Text, View, FlatList } from 'react-native';
import {SearchBar, CheckBox, ListItem} from "react-native-elements";
import Config from '../../Config';

class SearchUserItem extends Component {
    keyExtractor = (item, index) => index.toString()

    renderItem = ({ item }) => (
        <ListItem
            bottomDivider
            title={item.title}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    );

    render() {
        return (
            <FlatList
                keyExtractor={this.keyExtractor}
                data={UserList}
                renderItem={this.renderItem}
            />
        );
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
            title={item.title}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    );

    updateSearch = (searchText) => {
        this.setState({ searchText });
        alert(this.state.searchText);
        /*fetch('http://192.168.43.150:8080/v1/user' + '?userName=' + this.state.searchText)
            .then((response) => response.json())
            .then((responseJson) => {

                this.setState({
                    UserList: responseJson.movies,
                }, function(){

                });

            })
            .catch((error) => {
                alert(error)
            });*/
    };

    cancelSearch = (searchText) => {
        this.setState({
            showSearchItem: false,
            searchText: '',
        });
        alert(this.state.searchText);

    };

    render() {
        const { searchText } = this.state.searchText;
        if (this.state.UserList !== null) {
            return (
                <View>
                    <SearchBar
                        placeholder={'搜索'}
                        onChangeText={ this.updateSearch }
                        onClear={this.cancelSearch}
                        value={ this.state.searchText }
                    />
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
                    <SearchBar
                        placeholder={'搜索'}
                        onChangeText={ this.updateSearch }
                        value={ this.state.searchText }
                    />
                </View>
            )
        }
    }
}
