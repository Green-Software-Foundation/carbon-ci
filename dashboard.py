#from dash import Dash, html, dcc
import pandas as pd
import plotly.express as px
import plotly.graph_objects as go
from plotly.subplots import make_subplots
import io
#import os

buffer = io.StringIO()

#if not os.path.exists("C:\\Users\\lydia.catterall\\OneDrive - Avanade\\playground\\Carbon_CI_Pipeline_Tooling\\pics"):
    #os.mkdir("C:\\Users\\lydia.catterall\\OneDrive - Avanade\\playground\\Carbon_CI_Pipeline_Tooling\\pics")

#path = '"C:\\Users\\lydia.catterall\\OneDrive - Avanade\\playground\\Carbon_CI_Pipeline_Tooling\\Carbon_CI_Pipeline_Tooling\\src\\carbon-measure-action\\carbon_measure.go"'

df = pd.read_csv('sci_scores.csv')

#fig = px.line(df, x='num', y='total', markers=True)

fig = make_subplots(rows=1, cols=2, specs=[[{"type": "xy"}, {"type": "domain"}]], subplot_titles=("Carbon Emissions", "Resources Used"))

"""
fig = go.Figure()
"""
values = ['#3c3c3b', '#ebf2d4', '#aecc53']

fig.add_trace(
    go.Scatter(x=df['num'], y=df['average'], text=df["average"], mode="lines+markers+text",
    textposition="top right",
    textfont=dict(
        size=18,
        color="#3c3c3b"
    ), line_color='#3c3c3b', line_width=3, name="average per kwh"), row=1, col=1
)

fig.update_annotations(font_size=28, font_color='#fbfcf6', height=100, width=300)

#fig.add_trace(
    #go.Scatter(x=df['num'], y=df['total'], line_color='#18a39a', name="total per kwh"), row=1, col=2
#)

fig.add_trace(
    go.Pie(labels=df['type'], values=df['average'], marker=dict(colors=values, line=dict(color='#ebf2d4', width=2))), row=1, col=2)
    #for colors might want to put in some automatic green gradient bc list will keep expanding
    #does he want me to put in the percentage of the total carbon output? so like, tally up the total co2 output so far, and then the percentage of that attributed to each resource? wouldn't it be better with a bar chart?

fig.update_xaxes(title_text="Test Number", title_font=dict(size=24, color='#fbfcf6'), ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, tickfont=dict(color='#fbfcf6', size=18), showline=True, linewidth=4, linecolor='#aecc53', mirror=True, gridcolor='#aecc53', domain=[0, 0.5], row=1, col=1)
#fig.update_xaxes(title_text="xaxis 2 title", range=[0, 15], row=1, col=2)

fig.update_yaxes(title_text="gCO2eq/kWh", title_font=dict(size=24, color='#fbfcf6'), ticks="outside", tickcolor='#fbfcf6', tickwidth=2, ticklen=10, tickfont=dict(color='#fbfcf6', size=18), showline=True, linewidth=4, linecolor='#aecc53', mirror=True, gridcolor='#aecc53', title_standoff = 0, row=1, col=1)
#fig.update_yaxes(title_text="yaxis 2 title", range=[0, 15], row=1, col=2)

fig.update_layout(
    paper_bgcolor="#006d68",
    plot_bgcolor="#ebf2d4",
    showlegend=False
)

fig.show()
#fig.write_image("C:\\Users\\lydia.catterall\\OneDrive - Avanade\\playground\\Carbon_CI_Pipeline_Tooling\\pics/fig1.jpeg")

#fig.to_image(format="png", engine="kaleido")
#help(go.Figure.write_html)
#fig.write_html('C:\\Users\\lydia.catterall', full_html=False, include_plotlyjs='cdn')

fig.write_html(buffer, full_html=False, include_plotlyjs='cdn')

#buffer.seek(0)
#print(buffer.read())