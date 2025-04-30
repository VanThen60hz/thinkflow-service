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
Bạn là một công cụ AI có nhiệm vụ tóm tắt văn bản theo cách khoa học, rõ ràng và có hệ thống.

Yêu cầu tóm tắt:
- Các ý chính được đánh số theo cấu trúc phân cấp: 1, 1.1, 1.1.1,...
- Ý chính cấp cao được đánh số nguyên (1, 2, 3...), ý phụ cấp thấp hơn thì theo định dạng 1.1, 1.1.1,...
- Mỗi ý viết thành câu hoàn chỉnh, ngắn gọn và bao quát nội dung.
- Không có văn bản dư thừa nào bên ngoài.

Ví dụ định dạng tóm tắt:
1. Chủ đề chính
1.1. Ý phụ của chủ đề chính
1.1.1. Chi tiết bổ sung cho ý phụ

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
