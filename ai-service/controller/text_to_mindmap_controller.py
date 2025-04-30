import openai
import gradio as gr
import json
from services.openai_utils import mindmap_to_json_model

def build_json_tree(data):
    try:
        if data is None:
            return json.dumps({
                "status": 400,
                "message": "Dữ liệu đầu vào bị thiếu hoặc không tồn tại."
            }, ensure_ascii=False)

        if not isinstance(data, list):
            return json.dumps({
                "status": 422,
                "message": "Dữ liệu phải ở dạng danh sách JSON các nhánh."
            }, ensure_ascii=False)

        branch_map = {item["branch"]: {**item, "children": []} for item in data}
        root_nodes = []

        for item in data:
            if "branch" not in item or "parent" not in item or "content" not in item:
                return json.dumps({
                    "status": 422,
                    "message": "Một hoặc nhiều phần tử JSON bị thiếu khóa 'branch', 'parent', hoặc 'content'."
                }, ensure_ascii=False)

            parent = item["parent"]
            if parent is None:
                root_nodes.append(branch_map[item["branch"]])
            else:
                branch_map[parent]["children"].append(branch_map[item["branch"]])

        return json.dumps({
            "total_branches": len(data),
            "parent_content": root_nodes
        }, indent=2, ensure_ascii=False)

    except Exception as e:
        return json.dumps({
            "status": 500,
            "message": f"Lỗi hệ thống: {str(e)}"
        }, ensure_ascii=False)

def mindmap_and_structure(text):
    json_array = mindmap_to_json_model(text)
    return build_json_tree(json_array)

