import React, {Component} from "react";
import {View, Text, Alert, StyleSheet} from 'react-native';
import HTTP from "../../Network/Network";
import Config from "../../Config";
import {Avatar, ListItem} from "react-native-elements";
import {NavigationActions} from "react-navigation";

export default class UserInfoForOthersScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            userID: -1,
            userName: '',
            avatarID: '',
        };
        this.params = this.props.navigation.state.params;
        //console.warn(this.params);
        this.userID = this.params ? this.params.userID : null;
        //console.warn(this.userID);
        this.fetchData = this.fetchData.bind(this);
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        let obj = {  };
        //console.warn(obj);
        HTTP.get(('/user/' + this.userID), obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    userID: response.userID,
                    userName: response.userName,
                    avatarID: response.avatarID,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = ({navigation}) => ({
        headerTitle: (<Text numberOfLines={1} style={{flex:1, color: '#298BFF', fontSize: 23}}>{navigation.state.params.userName}</Text>)
    });

    render() {
        if (!this.state.loaded) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else {
            return (
                <View style={{backgroundColor: '#EFEFF5'}}>
                    <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
                    <ListItem
                        topDivider
                        bottomDivider
                        leftAvatar={() => {
                            if(this.state.avatarID === 'to_be_changed') {
                                return (
                                    <Avatar rounded size='large' source={require('../../assets/images/NotLogin.jpg')}/>
                                )
                            } else {
                                return (
                                    <Avatar rounded size='large' source={{uri: (Config.fetchPrefix + 'file/' + this.state.avatarID)}}/>
                                )
                            }
                        }}
                        title={this.state.userName}
                        titleStyle={{ color: 'black', fontSize: 25 }}
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
                        title={'TA发布的出售信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => {
                            let userID = this.userID;
                            let userName = this.state.userName;
                            //console.warn(userID);
                            this.props.navigation.push('SellInfoForOthers', { userID, userName });
                        }}
                    />
                    <ListItem
                        bottomDivider
                        rightIcon={<Avatar
                            rounded size='small'
                            source={require('../../assets/icons/right.png')}
                            overlayContainerStyle={{backgroundColor: 'white'}}
                        />}
                        title={'TA发布的求购信息'}
                        titleStyle={{ color: 'black', fontSize: 17 }}
                        onPress={() => {
                            let userID = this.userID;
                            let userName = this.state.userName;
                            this.props.navigation.push('BuyInfoForOthers', {userID, userName})
                        }}
                    />
                    <ListItem
                        containerStyle={{height: 1000, backgroundColor: '#EFEFF5'}}
                    />
                </View>
            )
        }
    }
}

let styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#EFEFF5"
    },
    subtitleView: {
        //flexDirection: 'row',
        paddingLeft: 10,
        paddingTop: 5
    },
    ratingText: {
        color: '#042A2B',
        fontSize: 18,
    },
    headerText: {
        color: '#5EB1BF',
        fontSize: 18,
    },
    goodType: {
        fontSize: 23,
        fontWeight: 'bold',
    }
});
