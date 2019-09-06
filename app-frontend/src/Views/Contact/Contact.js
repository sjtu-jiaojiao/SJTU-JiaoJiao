import React, {Component} from 'react';
import {
    View, Text, StyleSheet, FlatList, Alert,
    Image, TouchableOpacity, NativeModules, Dimensions
} from 'react-native';
import {Avatar, ListItem} from "react-native-elements";
import Config, {isLogin} from "../../Config";

export default class ContactScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            refreshing: false,
            contactPerson: [
                {"nickname":"Alex Black","location":"London","avatar":"1","header":"A"},
                {"nickname":"Alex Proti","location":"Moscow","avatar":"5","header":"A"},
                {"nickname":"Andrew Smith","location":"Kiev","avatar":"3","header":"A"},
                {"nickname":"Ann Ryder","location":"Kiev","avatar":"7","header":"A"},
                {"nickname":"Daniel Ricci","location":"Kiev","avatar":"8","header":"D"},
                {"nickname":"Ivan Ivanov","location":"Kiev","avatar":"3","header":"I"},
                {"nickname":"Kate Lebedeva","location":"Odessa","avatar":"6","header":"K"},
                {"nickname":"Kate Shy","location":"Kiev","avatar":"10","header":"K"},
                {"nickname":"Michael Fold","location":"Praha","avatar":"1","header":"M"},
                {"nickname":"Nadya Lovin","location":"Moscow","avatar":"2","header":"N"},
                {"nickname":"Oleg Price","location":"Odessa","avatar":"4","header":"O"},
                {"nickname":"Oleg Ryzhkov","location":"Kiev","avatar":"5","header":"O"},
                {"nickname":"Olga Blare","location":"Praha","avatar":"9","header":"O"},
                {"nickname":"Svetlana Kot","location":"Milan","avatar":"10","header":"S"}
            ]
        };
    }

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>我的消息</Text>),
    };

    componentDidMount() {
        /*this.setState({
            refreshing: true
        });
        this.props.fetchContacts()
            .then(_ => {
                this.setState({
                    refreshing: false
                })
            })*/
    }

    getRemoteAvatar = (id) => {
        return `https://loremflickr.com/70/70/people?lock=${id}`
    };

    keyExtractor = (item, index) => index.toString();

    renderItem = ({ item }) => {
        return (
            <View>
                <ListItem
                    topDivider
                    bottomDivider
                    containerStyle={styles.listItem}
                    leftAvatar={{ source: { uri: this.getRemoteAvatar(item.avatar) } }}
                    rightIcon={<Avatar
                        rounded size='small'
                        source={require('../../assets/icons/right.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />}
                    title={item.nickname}
                    subtitle={item.location}
                    subtitleStyle={styles.subtitleStyle}
                    onPress={() => { this.props.navigation.navigate('Message', { user: item }) }}
                />
            </View>
        )
    };

    render() {
        if (!isLogin()) {
            Alert.alert(
                '未登录',
                '无法在未登录状态下进入消息界面，是否切换至登录界面？',
                [
                    {
                        text: '取消',
                        onPress: () => {
                            this.props.navigation.navigate('Home');
                        },
                        style: 'cancel',
                    },
                    {
                        text: '确定', onPress: () => {
                            this.props.navigation.navigate('Login');
                        }
                    },
                ],
                {cancelable: false},
            );
            return (<View/>);
        } else {
            Config.isContactRender = true;
            return (
                <FlatList
                    style={styles.container}
                    contentContainerStyle={styles.contentContainerStyle}
                    keyExtractor={this.keyExtractor}
                    data={this.state.contactPerson}
                    renderItem={this.renderItem}
                    onRefresh={() => null}
                    refreshing={this.state.refreshing}
                />
            );
        }
    }
}

let styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: "#EFEFF5"
    },
    contentContainerStyle: {
        paddingBottom: 20,
    },
    listItem: {
        paddingTop: 8,
        paddingBottom: 8
    },
    subtitleStyle: {
        fontSize: 14,
        color: '#858585'
    }
});

