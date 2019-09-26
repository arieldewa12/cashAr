demo = {
    initPickColor: function() {
      $('.pick-class-label').click(function() {
        var new_class = $(this).attr('new-class');
        var old_class = $('#display-buttons').attr('data-class');
        var display_div = $('#display-buttons');
        if (display_div.length) {
          var display_buttons = display_div.find('.btn');
          display_buttons.removeClass(old_class);
          display_buttons.addClass(new_class);
          display_div.attr('data-class', new_class);
        }
      });
    },

  
    initChartsPages: function() {
      chartColor = "#FFFFFF";
  
      var speedCanvas = document.getElementById("speedChart");
  
      var dataFirst = {
        data: [0, 19, 15, 20, 30, 40, 40, 50, 25, 30, 50, 70],
        fill: false,
        borderColor: '#fbc658',
        backgroundColor: 'transparent',
        pointBorderColor: '#fbc658',
        pointRadius: 4,
        pointHoverRadius: 4,
        pointBorderWidth: 8,
      };
  
      var dataSecond = {
        data: [0, 5, 10, 12, 20, 27, 30, 34, 42, 45, 55, 63],
        fill: false,
        borderColor: '#51CACF',
        backgroundColor: 'transparent',
        pointBorderColor: '#51CACF',
        pointRadius: 4,
        pointHoverRadius: 4,
        pointBorderWidth: 8
      };

      var dataThird = {
        data: [1, 6, 11, 14, 21, 23, 33, 32, 40, 40, 50, 60],
        fill: false,
        borderColor: '#6bd098',
        backgroundColor: 'transparent',
        pointBorderColor: '#6bd098',
        pointRadius: 4,
        pointHoverRadius: 4,
        pointBorderWidth: 8
      };

      var dataFour = {
        data: [10, 5, 10, 12, 30, 37, 30, 14, 22, 15, 55, 13],
        fill: false,
        borderColor: '#ef8157',
        backgroundColor: 'transparent',
        pointBorderColor: '#ef8157',
        pointRadius: 4,
        pointHoverRadius: 4,
        pointBorderWidth: 8
      };
  
      var speedData = {
        labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
        datasets: [dataFirst, dataSecond,dataThird,dataFour]
      };
  
      var chartOptions = {
        legend: {
          display: false,
          position: 'top'
        }
      };
  
      var lineChart = new Chart(speedCanvas, {
        type: 'line',
        hover: false,
        data: speedData,
        options: chartOptions
      });
    },    
  
    showNotification: function(from, align) {
      color = 'primary';
  
      $.notify({
        icon: "nc-icon nc-bell-55",
        message: "Welcome to <b>Paper Dashboard</b> - a beautiful bootstrap dashboard for every web developer."
  
      }, {
        type: color,
        timer: 8000,
        placement: {
          from: from,
          align: align
        }
      });
    }
  
  };