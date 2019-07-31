import { Component, OnInit } from '@angular/core';
import { EChartOption} from 'echarts';
import * as echarts from 'echarts/lib/echarts';
import { InfoService } from '../info.service';
import { buyInfo, sellInfo } from 'src/app/entity/info';
import {prepareBoxplotData} from 'echarts/extension/dataTool';
const name = ['LJH', 'WXZ', 'ZWJ', 'KHQ', 'MZD', 'ZEL', 'JZM', 'HJT', 'TRUMP',
'LJH2', 'WXZ2', 'ZWJ2', 'KHQ2', 'MZD2', 'ZEL2', 'JZM2', 'HJT2', 'TRUMP2',
'LJH3', 'WXZ3', 'ZWJ3', 'KHQ3', 'MZD3', 'ZEL3', 'JZM3', 'HJT3', 'TRUMP3'];
@Component({
  selector: 'app-info-statistic',
  templateUrl: './info-statistic.component.html',
  styleUrls: ['./info-statistic.component.css']
})
export class InfoStatisticComponent implements OnInit {
  fdgoption: EChartOption;
  cldoption: any;
  tsoption: any;
  lqoption: any;
  goodoption: any;
  bi: buyInfo[];
  si: sellInfo[];
  constructor(private is: InfoService) { }
  ngOnInit() {
    this.is.getBuyInfos().subscribe(
        e => {this.bi = e.buyInfo;
            this.is.getSellInfos().subscribe(
                e => {
                    this.si = e.sellInfo;
                    this.cld();
                    this.fdg();
                    this.ts();
                    this.lq();
                    this.good();
                    this.getMoreBuy(100);
                    this.getMoreSell(100);
                }
  );})
  }
  getMoreBuy(offset){
      if(!(this.bi.length%100))
        this.is.getBuyInfos(null,null,null,null,offset).subscribe(
            e => {
                this.bi=this.bi.concat(e.buyInfo);
                this.lq();
                this.getMoreBuy(offset+100);
            }
        );
  }
  getMoreSell(offset){
      console.log(this.si.length);
      if(!(this.si.length%100))
        this.is.getSellInfos(null,null,null,null,offset).subscribe(
            e => {
                this.si= this.si.concat(e.sellInfo);
                this.lq();
                this.getMoreSell(offset+100);
            }
        );
  }
  good() {
    var data = prepareBoxplotData([
    [850, 740, 900, 1070, 930, 850, 950, 980, 980, 880, 1000, 980, 930, 650, 760, 810, 1000, 1000, 960, 960],
    [960, 940, 960, 940, 880, 800, 850, 880, 900, 840, 830, 790, 810, 880, 880, 830, 800, 790, 760, 800],
    [880, 880, 880, 860, 720, 720, 620, 860, 970, 950, 880, 910, 850, 870, 840, 840, 850, 840, 840, 840],
    [890, 810, 810, 820, 800, 770, 760, 740, 750, 760, 910, 920, 890, 860, 880, 720, 840, 850, 850, 780],
    [890, 840, 780, 810, 760, 810, 790, 810, 820, 850, 870, 870, 810, 740, 810, 940, 950, 800, 810, 870]
]);
this.goodoption = {
    backgroundColor: '#01193d',
    title: [
        {
            text: 'Price Trend',
            left: 'center',
        }
    ],
    tooltip: {
        trigger: 'item',
        axisPointer: {
            type: 'shadow'
        }
    },
    dataZoom: [
        {
            show: true,
            realtime: true,
            start: 0,
            end: 40
        },
        {
            type: 'inside',
            realtime: true,
            start: 0,
            end: 40
        }
    ],
    grid: {
        left: '10%',
        right: '10%',
        bottom: '15%'
    },
    xAxis: {
        type: 'category',
        data: data.axisData,
        boundaryGap: true,
        nameGap: 30,
        axisLine: {onZero: false},
        splitArea: {
            show: false
        },
        axisLabel: {
            formatter: 'tag {value}'
        },
        splitLine: {
            show: false
        },
    },
    yAxis: {
        type: 'value',
        name: 'yuan'
    },
    series: [
        {
            name: 'boxplot',
            type: 'boxplot',
            data: data.boxData,
            tooltip: {
                formatter: function (param) {
                    return [
                        'tag ' + param.name + ': ',
                        'upper: ' + param.data[5],
                        'Q3: ' + param.data[4],
                        'median: ' + param.data[3],
                        'Q1: ' + param.data[2],
                        'lower: ' + param.data[1]
                    ].join('<br/>');
                }
            }
        },
        {
            name: 'outlier',
            type: 'scatter',
            data: data.outliers
        }
    ]
};
  }
  ts() {
    this.tsoption = {
      backgroundColor: '#01193d',
      title: {
          text: 'Transaction Calendar',
          left: 'center'
      },
      tooltip : {
          trigger: 'item'
      },
      grid: {
        containLabel: true,
      },
      calendar: [{
          left: 80,
          range: ['2019'],
          splitLine: {
              show: true,
              lineStyle: {
                  color: '#000',
                  width: 2,
                  type: 'solid'
              }
          },
          width: '80%',
          height: '80%',
          dayLabel: {
            textStyle: {
                color: '#fff'
            }
          },
          monthLabel: {
            textStyle: {
                color: '#fff'
            }
          },
          yearLabel: {
              formatter: '{start}',
              textStyle: {
                  color: '#fff'
              }
          },
          itemStyle: {
              normal: {
                  color: '#323c48',
                  borderWidth: 1,
                  borderColor: '#111'
              }
          }
      }],
          series: [{
            type: 'effectScatter',
            coordinateSystem: 'calendar',
            symbolSize: (val) => {
                return val[1] / 40;
            },
            showEffectOn: 'render',
            rippleEffect: {
                brushType: 'stroke'
            },
            hoverAnimation: true,
            itemStyle: {
                    color: '#f4e925',
                    shadowBlur: 10,
                    shadowColor: '#333'
            },
              data: [['2019-01-02', 900], ['2019-01-03', 877], ['2019-01-04', 699], ['2019-01-07', 200], ['2019-01-10', 100],
              ['2019-01-10', 430], ['2019-02-01', 250], ['2019-02-10', 430],
              ['2019-03-10', 430], ['2019-04-01', 250], ['2019-05-10', 430],
              ['2019-08-11', 430], ['2019-07-04', 250], ['2019-03-11', 430],
              ['2019-09-23', 430], ['2019-06-01', 250], ['2019-12-12', 430]]
          }]
        };
  }
  cld() {
    this.cldoption = {
    backgroundColor: '#01193d',
      title: {
      text: 'Label WordCloud',
      left: 'center',
  },
  tooltip: {},
  series: [{
      type: 'wordCloud',
      shape: 'circle',

      left: 'center',
      top: 'center',
      sizeRange: [12, 30],
      rotationRange: [-90, 90],
      rotationStep: 45,
      gridSize: 8,
      drawOutOfBound: false,
      textStyle: {
          normal: {
              fontFamily: 'sans-serif',
              fontWeight: 'bold',
              color: () => {
                  // Random color
                  return 'rgb(' + [
                      Math.round(Math.random() * 250),
                      Math.round(Math.random() * 250),
                      Math.round(Math.random() * 250)
                  ].join(',') + ')';
              }
          },
          emphasis: {
              shadowBlur: 10,
              shadowColor: '#333'
          }
      },
      data: name.map( node => {return { name: node,
      value: Math.round(Math.random() * 1000),
      textStyle: {
          normal: {},
          emphasis: {}
      }
  }; }
  )
  }]
};
  }
  fdg() {
    const E = name.map(node => {
      const e = Math.floor(Math.random() * 10);
      return {source: node, target: name[Math.floor(Math.random() * 27)]
      , value: e, lineStyle: {
        width: e * 0.3
      }};
  }).filter(node => node.source !== node.target);
    const V = name.map(node => {
      const v = E.filter(n => n.source === node || n.target === node ).reduce((total, currentValue, currentIndex, arr) => {
      return total + currentValue.value;
  }, 0);
      return { name: node,
      itemStyle: {
        color: '#60acfc'
      }, value : v, symbolSize: v,
  draggable: true};
});
    this.fdgoption = {
    backgroundColor: '#01193d',
      title: {
          text: 'Transaction Network',
          left: 'center',
      },
      tooltip: {},
      animationDurationUpdate: 1500,
      animationEasingUpdate: 'quinticInOut',
      series : [
          {
              type: 'graph',
              layout: 'force',
              data: V, edges: E,
        label: {
          emphasis: {
              position: 'right',
              show: true
          }
      },
      force: {
        repulsion : 100
      },
      roam: true,
      focusNodeAdjacency: true,
      lineStyle: {
          normal: {
              color :'source',
              type : 'solid',
              width: 0.5,
              curveness: 0.2,
              opacity: 0.7
          }
      }
  }
          ]
      };
  }
  fmt(t: Date) {
    return [t.getFullYear(), t.getMonth() + 1, t.getDate()].join('/');
  }
  lq() {
    let bd =new Map();
    this.bi.forEach( e => {
        if(!e||e.releaseTime<0)return;
       const str = this.fmt(new Date(e.releaseTime*1000));
       if(str in bd)
        bd[str]+=1;
        else 
        bd[str]=1;
    });
    let sd =new Map();
    this.si.forEach( e => {
        if(!e||e.releaseTime<0)return;
       const str = this.fmt(new Date(e.releaseTime*1000));
       if(str in sd)
        sd[str]+=1;
        else 
        sd[str]=1;
    });
    let bdata = [];
    for(let i in bd){
        bdata.push([i,bd[i]]);
    }

    let sdata = [];
    for(let i in sd){
        sdata.push([i,sd[i]]);
    }
    this.lqoption = {
    backgroundColor: '#01193d',
    title : {
        text: 'Transaction Trend',
        x: 'center',
        align: 'right'
    },
    grid: {
        bottom: 80
    },
    tooltip : {
        trigger: 'axis',
        axisPointer: {
            type: 'cross',
            animation: false,
            label: {
                backgroundColor: '#505765'
            }
        }
    },
    legend: {
        data: ['购买', '出售'],
        x: 'left'
    },
    dataZoom: [
        {
            show: true,
            realtime: true,
            start: 0,
            end: 100
        },
        {
            type: 'inside',
            realtime: true,
            start: 0,
            end: 100
        }
    ],
    xAxis : [
        {
            type : 'time',
            boundaryGap : false,
            axisLine: {onZero: false},
        }],
    yAxis: [
        {
            name: '新增购买信息',
            type: 'value',
            max: (value) => value.max*3/2
        },
        {
            name: '新增出售信息',
            nameLocation: 'start',
            type: 'value',
            inverse: true,
            max: (value) => value.max*3/2
        }
    ],
    series: [
        {
            name: '购买',
            type: 'line',
            animation: false,
            smooth: true,
            showAllSymbol: true,
            symbol: 'circle',
            symbolSize: 6,
            areaStyle: {
                normal: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                        offset: 0,
                        color: 'rgba(137, 189, 27, 0.9)'
                    }, {
                        offset: 0.8,
                        color: 'rgba(137, 189, 27, 0)'
                    }], false),
                    shadowColor: 'rgba(0, 0, 0, 0.1)',
                    shadowBlur: 10
                }
            },
            itemStyle: {
                normal: {
                    color: 'rgb(137,189,27)',
                    borderColor: 'rgba(137,189,2,0.27)',
                    borderWidth: 12
                }
            },
            lineStyle: {
                width: 1
            },
            data: bdata
        },
        {
            name: '出售',
            type: 'line',
            yAxisIndex: 1,
            animation: false,
            smooth: true,
            showAllSymbol: true,
            symbol: 'circle',
            symbolSize: 6,
            areaStyle: {
                normal: {
                    color: new echarts.graphic.LinearGradient(0, 1, 0,0 , [{
                        offset: 0,
                        color: 'rgba(219, 50, 51, 0.9)'
                    }, {
                        offset: 0.8,
                        color: 'rgba(219, 50, 51, 0)'
                    }], false),
                    shadowColor: 'rgba(0, 0, 0, 0.1)',
                    shadowBlur: 10
                }
            },
            itemStyle: {
                normal: {
                    color: 'rgb(219,50,51)',
                    borderColor: 'rgba(219,50,51,0.2)',
                    borderWidth: 12
                }
            },
            lineStyle: {
                width: 1
            },
            data: sdata
        }
    ]
};

  }
}
