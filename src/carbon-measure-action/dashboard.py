"""
This file uses Plotly to construct a dashboard figure with two subplots:
    - A connected scatter plot keeping track of the carbon intensity of one resource.
    - A pie chart showing the percentage of the carbon emissions attributed to each resource. 

To compare the average carbon emissions of each resource side by side instead of viewing the percentage attributed to each resource, follow the comments to use a bar chart instead of a pie chart.
"""

import pandas as pd
import plotly.express as px
import plotly.graph_objects as go
from plotly.subplots import make_subplots

# The csv currently has sample values with columns for the build number, the average carbon intensity, and the total carbon intensity.
df = pd.read_csv('sample_scores.csv')

# For pie chart, replace the second type with "xy"
fig = make_subplots(rows=1, cols=2, specs=[[{"type": "xy"}, {"type": "domain"}]], subplot_titles=("Current Resource SCI", "SCI of Each Resource"))

# Constructs a connected scatter plot
fig.add_trace(
    go.Scatter(x=df['num'], y=df['average'], text=df["average"], mode="lines+markers+text",
    textposition="top right",
    textfont=dict(
        size=18,
        color="#3c3c3b"
    ), line_color='#3c3c3b', line_width=3, name="average per kwh"), row=1, col=1
)

fig.update_xaxes(title_text="Test Number", 
    title_font=dict(size=24, color='#fbfcf6'), 
    ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, 
    tickfont=dict(color='#fbfcf6', size=18), 
    gridcolor='#aecc53', 
    row=1, col=1
)

fig.update_yaxes(title_text="gCO2eq/kWh", 
    title_font=dict(size=24, color='#fbfcf6'), 
    ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, 
    tickfont=dict(color='#fbfcf6', size=18), 
    gridcolor='#aecc53', 
    title_standoff = 0, 
    row=1, col=1
)

# Constructs a pie chart (should eventually draw labels and values from a csv)
colorlist= 5 * px.colors.sequential.Emrld
fig.add_trace(
    go.Pie(labels=['rsrc_1', 'rsrc_2', 'rsrc_3', 'rsrc_4'], values=[5,2,3,4], 
    marker=dict(colors=colorlist, line=dict(color='#ebf2d4', width=2)), 
    hoverinfo='label+percent', textinfo='label+percent', textfont_size=20, 
    name="resources used"), 
    row=1, col=2
)

# Uncomment the following lines to use a bar chart instead of a pie chart
"""
# Currently has placeholder values for x and y axes
fig.add_trace(go.Bar(x=['a', 'b', 'c'], y=[1, 2, 3], marker_color='#3c3c3b', name='resources used'), row=1, col=2)

fig.update_xaxes(title_text="Resource Type", 
    title_font=dict(size=24, color='#fbfcf6'), 
    ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, 
    tickfont=dict(color='#fbfcf6', size=18), 
    gridcolor='#aecc53', 
    row=1, col=2
)

fig.update_yaxes(title_text="gCO2eq/kWh", 
    title_font=dict(size=24, color='#fbfcf6'), 
    ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, 
    tickfont=dict(color='#fbfcf6', size=18), 
    gridcolor='#aecc53', 
    title_standoff = 0, 
    row=1, col=2
)
"""
# Specifications for the subplot titles
fig.update_annotations(font_size=28, font_color='#fbfcf6', height=100, width=300)

fig.update_layout(
    paper_bgcolor="#3c3c3b",
    plot_bgcolor="#ebf2d4",
    showlegend=False,
    #barmode='group' # Uncomment when using a bar chart instead of pie chart
)

# Displays the figure in a web browser each time the program is run
fig.show()

#Export the figure as a static image: https://plotly.com/python/static-image-export/
#Export the figure as an interactive html file: https://plotly.com/python/interactive-html-export/