import React, { Component } from 'react';
import { Text, View, FlatList } from 'react-native';
import {SearchBar, CheckBox, ListItem} from "react-native-elements";

const DemoUser = [
    {
        userName: 'chiangel',
    },
    {
        userName: 'wxz',
    },
    {
        userName: 'Jboss',
    },
    {
        userName: 'Boy',
    },
    {
        userName: 'a',
    },
    {
        userName: 'b',
    },
    {
        userName: 'vc',
    },
    {
        userName: 'd',
    },
    {
        userName: 'e',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
    {
        userName: 'y',
    },
];

class SearchUserItem extends Component {
    keyExtractor = (item, index) => index.toString()

    renderItem = ({ item }) => (
        <ListItem
            bottomDivider
            title={item.userName}
            titleStyle={{ color: 'black', fontSize: 17 }}
        />
    )

    render() {
        return (
            <FlatList
                keyExtractor={this.keyExtractor}
                data={DemoUser}
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
        }
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23}}></Text>)
    };

    updateSearch = (searchText) => {
        this.setState({ searchText });
        this.setState(previousState => (
            { showSearchItem: true }
        ));
    };

    cancelSearch = (searchText) => {
        this.setState({ searchText });
        this.setState(previousState => (
            {
                showSearchItem: false,
                searchText: '',
            }
        ));
    };

    render() {
        const { searchText } = this.state.searchText;
        if (this.state.showSearchItem) {
            return (
                <View>
                    <SearchBar
                        placeholder={'搜索'}
                        onChangeText={ this.updateSearch }
                        onClear={this.cancelSearch}
                        value={ this.state.searchText }
                    />
                    <SearchUserItem/>
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