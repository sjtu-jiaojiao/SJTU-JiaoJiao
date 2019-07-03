import { Component, OnInit } from '@angular/core';
import { EChartOption } from 'echarts';
import * as $ from 'jquery';
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
  constructor() { }
  ngOnInit() {
    this.cld();
    this.fdg();
    this.ts();
    this.lq();
    this.good();
  }
  onChartInit(ec) {
      ec.resize({height: 620});
    }
  good() {
    const jsdata1: any[][] = [['2019-01-01', 18.5, 3], ['2019-01-02', 12.5, 3], ['2019-01-01', 42.5, 5], ['2019-01-03', 65, 5],
    ['2019-01-05', 23.5, 5], ['2019-01-04', 18, 5], ['2019-01-02', 18, 4]];
    const jsdata2: any[][] = [['2019-01-01', 32.5, 2], ['2019-01-02', 122.5, 4], ['2019-01-01', 41.5, 5], ['2019-01-07', 75, 5],
    ['2019-01-03', 21.5, 5], ['2019-01-04', 177, 5], ['2019-01-02', 66, 5]];
    const jsdata3: any[][]  = [['2019-01-01', 11], ['2019-01-06', 12.5, 4], ['2019-01-01', 4.5, 3], ['2019-01-03', 15, 3],
    ['2019-01-03', 11.5, 5], ['2019-01-04', 17, 3], ['2019-01-02', 16, 5]];
    const itemstyle = {
    normal: {
        opacity: 0.8,
        shadowBlur: 10,
        shadowOffsetX: 0,
        shadowOffsetY: 0,
        shadowColor: 'rgba(0, 0, 0, 0.3)'
    }
};
    this.goodoption = {
        title: {
            text: 'Price Trend',
        },
        color: [
            '#dd4444', '#fec42c', '#80F1BE'
        ],
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'cross'
            }
        },
        legend: {
            y: 'top',
            data: ['Book', 'Shoe', 'Mouse'],
        },
        xAxis: {
            type: 'category',
        },
        yAxis: {
            type: 'value',
            min: 0,
            max: 'datamax',
            name: 'Price',
            splitLine: {
                show: true
            }
        },
        series: [{
            name: 'Book',
            type: 'scatter',
            symbolSize: t => 3 * t[2],
            itemStyle: itemstyle,
            data: jsdata1
        }, {
            name: 'Shoe',
            type: 'scatter',
            symbolSize: t => 3 * t[2],
            itemStyle: itemstyle,
            data: jsdata2
        },
        {
            name: 'Mouse',
            type: 'scatter',
            symbolSize: t => 3 * t[2],
            itemStyle: itemstyle,
            data: jsdata3
        }

    ]
    };
  }
  ts() {
    this.tsoption = {
      title: {
          text: 'Transaction Calendar'
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
      title: {
      text: 'Label WordCloud'
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
      title: {
          text: 'Transaction Network'
      },
      tooltip: {},
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
              width: 0.5,
              curveness: 0.3,
              opacity: 0.7
          }
      }
  }
          ]
      };
  }
  lq() {
    const date =  [
        '2009/6/12 2:00', '2009/6/12 3:00', '2009/6/12 4:00', '2009/6/12 5:00', '2009/6/12 6:00',
        '2009/6/12 7:00', '2009/6/12 8:00', '2009/6/12 9:00', '2009/6/12 10:00', '2009/6/12 11:00',
        '2009/6/12 12:00', '2009/6/12 13:00', '2009/6/12 14:00', '2009/6/12 15:00', '2009/6/12 16:00',
        '2009/6/12 17:00', '2009/6/12 18:00', '2009/6/12 19:00', '2009/6/12 20:00', '2009/6/12 21:00',
        '2009/6/12 22:00', '2009/6/12 23:00'].map( (str) => {
        return str.replace(' ', '\n'); }
        );
    this.lqoption = {
    title : {
        text: 'Transaction Trend',
        x: 'center',
        align: 'right'
    },
    grid: {
        bottom: 80
    },
    toolbox: {
        feature: {
            dataZoom: {
                yAxisIndex: 'none'
            },
            restore: {},
            saveAsImage: {}
        }
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
            start: 65,
            end: 85
        },
        {
            type: 'inside',
            realtime: true,
            start: 65,
            end: 85
        }
    ],
    xAxis : [
        {
            type : 'category',
            boundaryGap : false,
            axisLine: {onZero: false},
            data : date
        }],
    yAxis: [
        {
            name: '新增购买信息',
            type: 'value',
            max: 500
        },
        {
            name: '新增出售信息',
            nameLocation: 'start',
            max: 500,
            type: 'value',
            inverse: true
        }
    ],
    series: [
        {
            name: '购买',
            type: 'line',
            animation: false,
            areaStyle: {
            },
            lineStyle: {
                width: 1
            },
            data: date.map(_ => Math.floor(Math.random() * 50))
        },
        {
            name: '出售',
            type: 'line',
            yAxisIndex: 1,
            animation: false,
            areaStyle: {
            },
            lineStyle: {
                width: 1
            },
            data: date.map(_ => Math.floor(Math.random() * 50))
        }
    ]
};

  }
}
