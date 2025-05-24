# app.py
import gradio as gr
import requests
import pandas as pd

API_BASE = "http://localhost:8080/api/todos"

def list_todos():
    """バックエンドから TODO 一覧を取得して DataFrame で返す"""
    resp = requests.get(API_BASE)
    resp.raise_for_status()
    todos = resp.json()
    # pandas.DataFrame に整形
    df = pd.DataFrame(todos)
    if df.empty:
        return "現在、TODO は登録されていません。"
    return df

def create_todo(title: str, description: str):
    """タイトルと説明から TODO を作成し、完了後一覧を返す"""
    payload = {"title": title, "description": description}
    resp = requests.post(API_BASE, json=payload)
    resp.raise_for_status()
    return list_todos()

with gr.Blocks() as demo:
    gr.Markdown("## 📝 TODO 管理アプリ（Go + Gradio）")
    
    with gr.Row():
        with gr.Column(scale=1):
            title_input = gr.Textbox(label="タイトル", placeholder="やることを入力")
            desc_input  = gr.Textbox(label="説明", placeholder="詳細を入力")
            add_btn     = gr.Button("追加")
        with gr.Column(scale=2):
            todo_table = gr.Dataframe(headers=["id","title","description"], datatype=["number","str","str"], label="TODO 一覧")
    
    # ボタン押下で create→list を実行し、テーブルを更新
    add_btn.click(fn=create_todo, inputs=[title_input, desc_input], outputs=todo_table)
    # 起動時に一覧を表示
    demo.load(fn=list_todos, inputs=None, outputs=todo_table)

if __name__ == "__main__":
    demo.launch()
