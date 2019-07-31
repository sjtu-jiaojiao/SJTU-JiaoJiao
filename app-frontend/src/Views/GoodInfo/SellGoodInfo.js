import React, { Component } from 'react';
import {View, Text, StyleSheet, ScrollView} from 'react-native';
import HTTP from '../../Network/Network';
import {parseStatus, parseTimeStamp} from "../../Utils/ParseInfo";
import MyAvatar from '../../Components/MyAvatar';

class InfoPart extends Component {
    render() {
        let {item} = this.props;
        return (
            <View style={{marginTop: 20}}>
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}} selectable={true}>
                        <Text style={styles.headerText}>出售</Text>
                        <Text style={styles.headerText}>出售物品：</Text>
                        <Text style={styles.ratingText}>{item.goodName}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}}>
                        <Text style={styles.headerText}>具体描述：</Text>
                        <Text style={styles.ratingText}>{item.description}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}}>
                        <Text style={styles.headerText}>商品状态：</Text>
                        <Text style={styles.ratingText}>{parseStatus(item.status)}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}}>
                        <Text style={styles.headerText}>出售价格：</Text>
                        <Text style={styles.ratingText}>￥{item.price}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}}>
                        <Text style={styles.headerText}>发布时间：</Text>
                        <Text style={styles.ratingText}>{parseTimeStamp(item.releaseTime)}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}}>
                        <Text style={styles.headerText}>商品标签：</Text>
                        <Text style={styles.ratingText}>暂无</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
            </View>
        );
    }
}

class ImagePart extends Component {
    render() {
        let {contentID} = this.props;
        return (
            <View style={{marginTop: 20}}>
                <View style={{backgroundColor: 'white'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}} selectable={true}>
                        <Text style={styles.headerText}>图片部分{contentID}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
            </View>
        )
    }
}

class UserPart extends  Component {
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            userID: -1,
            userName: '',
            avatarID: '',
        };
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        let {userID} = this.props;
        //console.warn(obj);
        //console.warn("http://202.120.40.8:30711/v1"+'/sellInfo/' + this.sellInfoID);
        HTTP.get(('/user/' + userID), obj)
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

    render() {
        let {userID} = this.props;
        return (
            <View style={{}}>
                <View style={{backgroundColor: 'white', flexDirection: 'row'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}} selectable={true}>
                        <Text style={styles.headerText}>用户信息部分{userID}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
            </View>
        )
    }
}

export default class SellGoodInfoScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            status: -1,
            releaseTime: -1,
            validTime: -1,
            goodName: '',
            price: -1,
            description: '',
            contentID: '',
            userID: -1,
        };
        this.params = this.props.navigation.state.params;
        //console.warn(this.params)
        this.sellInfoID = this.params ? this.params.sellInfoID : null;
        //console.warn(this.sellInfoID);
        this.fetchData = this.fetchData.bind(this);
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        //console.warn(obj);
        //console.warn("http://202.120.40.8:30711/v1"+'/sellInfo/' + this.sellInfoID);
        HTTP.get(('/sellInfo/' + this.sellInfoID), obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    status: response.status,
                    releaseTime: response.releaseTime,
                    validTime: response.validTime,
                    goodName: response.goodName,
                    price: response.price,
                    description: response.description,
                    contentID: response.contentID,
                    userID: response.userID,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = ({navigation}) => ({
        headerTitle: (<Text numberOfLines={1} style={{flex:1, color: '#298BFF', fontSize: 23}}>{navigation.state.params.goodName}</Text>)
    });

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else {
            //console.warn(this.state.SellInfoList[0].releaseTime);
            //console.warn(TimeStamptoDate(this.state.SellInfoList[0].releaseTime));
            //console.warn(TimeStampNow());
            //console.warn(TimeStamptoDate(TimeStampNow()));
            return (
                <ScrollView>
                    <View style={{backgroundColor: '#EFEFF5'}}>
                        <InfoPart item={this.state}/>
                        <ImagePart contentID={this.state.contentID}/>
                        <UserPart userID={this.state.userID}/>
                    </View>
                </ScrollView>
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
        color: 'black',
        fontSize: 20,
    },
    headerText: {
        color: '#E53A40',
        fontSize: 20,
        fontWeight: 'bold',
    }
});
