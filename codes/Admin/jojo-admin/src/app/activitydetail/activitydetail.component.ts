import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-activitydetail',
  templateUrl: './activitydetail.component.html',
  styleUrls: ['./activitydetail.component.css']
})
export class ActivitydetailComponent implements OnInit {
  tsoption = {
    backgroundColor: '#01193d',
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
  constructor() { }

  ngOnInit() {
  }

}
