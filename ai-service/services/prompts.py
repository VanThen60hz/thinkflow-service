mindmap_prompt = """
Bạn là một công cụ AI có nhiệm vụ tóm tắt văn bản thành dạng cây phân nhánh đơn giản.
Trả kết quả dưới dạng JSON array gồm các đối tượng có cấu trúc:
- "branch": tên nhánh, ví dụ "branch_1"
- "parent": nếu là nhánh chính thì là null, nếu không thì chỉ rõ tên nhánh cha (ví dụ: "branch_1")
- "content": nội dung tóm tắt của nhánh đó

Chỉ trả JSON thuần, không có giải thích hay văn bản dư thừa nào bên ngoài.

Văn bản cần tóm tắt:
\"\"\"{text}\"\"\"
"""

summary_prompt = """
Bạn là một công cụ AI có nhiệm vụ tóm tắt văn bản một cách khoa học nhất (chia ra các ý chính và các ý phụ bổ sung cho nó).

Chỉ cần chỉnh lại những từ điện phương thành từ phổ thông, không có giải thích hay thêm văn bản dư thừa nào bên ngoài.

Văn bản cần tóm tắt: 
\"\"\"{text}\"\"\"
"""

audio_prompt = """
Bạn là một công cụ AI có nhiệm vụ điều chỉnh những từ địa phương được transcrips bị sai sao cho có nghĩa trong ngữ cảnh 
(ví dụ: Cùng họp bắt đầu -> cuộc họp bắt đầu, phát tài hiệu trước -> phát tài liệu trước, Mùi họp -> buổi họp, ...)

Chỉ cần chỉnh lại những từ điện phương thành từ phổ thông có cấu trúc tương tự nhất, không có giải thích hay thêm văn bản dư thừa nào bên ngoài.

Văn bản cần điều chỉnh: 
\"\"\"{text}\"\"\"
"""
