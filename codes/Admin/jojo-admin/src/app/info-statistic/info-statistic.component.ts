import { Component, OnInit } from '@angular/core';
import { EChartOption } from 'echarts';

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
  constructor() { }
  ngOnInit() {
    this.cld();
    this.fdg();
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
    this.fdgoption = {
      title: {
          text: 'Transaction Network'
      },
      tooltip: {},
      series : [
          {
              type: 'graph',
              layout: 'force',
              data: name.map(node => { return { name: node,
                itemStyle: {
                  color: '#60acfc'
            }, value : 1,
            draggable: true};
          }), edges: name.map(node => {return {source: node, target: name[Math.floor(Math.random() * 27)]
            , value: Math.floor(Math.random() * 10)};
        }).filter(node => node.source !== node.target),
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
}
