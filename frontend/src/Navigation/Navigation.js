import React, { Component } from 'React';
import { View, Text } from 'react-native';
import {
    createBottomTabNavigator,
    createStackNavigator,
    createAppContainer
} from "react-navigation";
import Icon from 'react-native-vector-icons/AntDesign';

import HomeScreen from '../Views/Home/Home'
import ReleaseScreen from '../Views/Release/Release'
import ContactScreen from '../Views/Contact/Contact'
import UserScreen from '../Views/User/User'
import MyBuyInfoScreen from "../Views/MyBuyInfo/MyBuyInfo"
import MySellInfoScreen from "../Views/MySellInfo/MySellInfo";
import MyHistoryInfoScreen from "../Views/MyHistoryInfo/MyHistoryInfo";
import UserInfoScreen from "../Views/UserInfo/UserInfo";
import BuyInfoScreen from "../Views/BuyInfo/BuyInfo";
import SellInfoScreen from "../Views/SellInfo/SellInfo";
import SearchScreen from "../Views/Search/Search";
import LoginScreen from "../Views/Login/Login";

class Test extends Component {
    render() {
        return (
            <View >
                <Text>
                    test page!
                </Text>
            </View>
        )
    }
}

const HomeStack = createStackNavigator({
    Home: { screen: HomeScreen },
    BuyInfo: { screen: BuyInfoScreen },
    SellInfo: { screen: SellInfoScreen },
    Search: { screen: SearchScreen },
});

const ReleaseStack = createStackNavigator({
    Release: { screen: ReleaseScreen },
    Test: { screen: Test},
});

const ContactStack = createStackNavigator({
    Contact: { screen: ContactScreen },
    Test: { screen: Test},
});

const UserStack = createStackNavigator({
    User: { screen: UserScreen },
    MyBuyInfo: { screen: MyBuyInfoScreen },
    MySellInfo: { screen: MySellInfoScreen },
    MyHistoryInfo: { screen: MyHistoryInfoScreen},
    UserInfo: { screen: UserInfoScreen},
    Login: { screen: LoginScreen },
});

const TabBar = createBottomTabNavigator({
    Home: { screen: HomeStack },
    Release: { screen: ReleaseStack },
    Contact: { screen: ContactStack },
    User: { screen: UserStack },
},{
    defaultNavigationOptions: ({ navigation }) => ({
        tabBarLabel: ({ focusd, tintColor }) => {
            const { routeName } = navigation.state;
            switch (routeName) {
                case 'Home':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>首页</Text>
                case 'Release':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>发布</Text>
                case 'Contact':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>消息</Text>
                case 'User':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>个人</Text>
            }
        },
        tabBarIcon: ({ focused, tintColor }) => {
            const { routeName } = navigation.state;
            let iconName;
            if (routeName === 'Home') {
                iconName = 'home'
            }
            else if (routeName === 'Release') {
                iconName = 'form'
            }
            else if (routeName === 'Contact') {
                iconName = 'message1'
            }
            else if (routeName === 'User') {
                iconName = 'user'
            }
            return <Icon name={iconName} size={25} color={tintColor} />;
        }
    })
});
/*
const AppStack = createStackNavigator({
    Tabs: TabBar,
    TestPage: Test,
}, {
    defaultNavigationOptions: {
        headerStyle: {
            backgroundColor: '#ff9800',
        },
        headerTintColor: '#fff',
    }
});*/

export default createAppContainer(TabBar);

export { Test };
