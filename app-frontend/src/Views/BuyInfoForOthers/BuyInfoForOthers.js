import React, { Component } from 'react';
import {FlatList, Text, View, StyleSheet, TouchableOpacity} from 'react-native';
import {Avatar, Button, ListItem, SearchBar, Icon} from "react-native-elements";
import Config from '../../Config';
import {TimeStamptoDate} from "../../Utils/TimeStamp";
import HTTP from '../../Network/Network';

export default class BuyInfoForOthersScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            BuyInfoList: null,
            loaded: false,
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
        let obj = { userID: this.userID };
        //console.warn(obj);
        HTTP.get('/buyInfo', obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    BuyInfoList: response.buyInfo,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = ({navigation}) => ({
        headerTitle: (<Text numberOfLines={1} style={{flex:1, color: '#298BFF', fontSize: 23}}>{navigation.state.params.userName}的出售信息</Text>)
    });

    keyExtractor = (item, index) => index.toString();

    parseStatus = (status) => {
        let Prefix = '商品状态：';
        switch (status) {
            case 1:
                return (Prefix + '出售中');
            case 2:
                return (Prefix + '已预约');
            case 3:
                return (Prefix + '已出售');
            case 4:
                return (Prefix + '已过期 (不再出售)');
        }
    };

    parseTimeStamp = (TimeStamp) => {
        let date = TimeStamptoDate(TimeStamp);
        return ('发布时间：' + date);
    };

    renderItem = ({ item }) => {
        let buyInfoID = item.buyInfoID;
        let header = ('求购：' + item.goodName);
        let infoType = 'buyInfo';
        return (
            <TouchableOpacity onPress={() => this.props.navigation.push('GoodInfo', { infoType, buyInfoID, header })}>
                <ListItem
                    bottomDivider
                    containerStyle={{height: 200}}
                    title={
                        <Text numberOfLines={1} style={{color: 'black', fontWeight: 'bold', fontSize: 20}}>出售物品：{item.goodName}</Text>
                    }
                    subtitle={
                        <View style={styles.subtitleView}>
                            <Text numberOfLines={1} style={styles.ratingText}>{this.parseStatus(item.status)}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>具体描述：{item.description}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>出售价格：￥{item.price}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>{this.parseTimeStamp(item.releaseTime)}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>有效时间：{item.validTime}</Text>
                            <Text numberOfLines={1} style={styles.ratingText}>物品标签：暂无</Text>
                        </View>
                    }
                />
            </TouchableOpacity>
        )
    };

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else if (this.state.BuyInfoList === undefined) {
            return (
                <View style={styles.container}>
                    <Text>TA暂时没有发布任何出售信息</Text>
                </View>
            )
        }
        else {
            return (
                <View>
                    <FlatList
                        keyExtractor={this.keyExtractor}
                        data={this.state.BuyInfoList}
                        renderItem={this.renderItem}
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
        paddingLeft: 10,
        color: 'grey',
        fontSize: 16,
    },
});
