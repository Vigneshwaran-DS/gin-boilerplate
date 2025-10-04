#!/bin/bash

# API 测试脚本 - Bash 版本
# 使用方法: ./scripts/api-test.sh

BASE_URL="http://localhost:8080/api/v1"
TOKEN=""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}=====================================${NC}"
echo -e "${CYAN}API 测试脚本${NC}"
echo -e "${CYAN}=====================================${NC}"
echo ""

# 1. 健康检查
echo -e "${YELLOW}1. 健康检查...${NC}"
response=$(curl -s -X GET "$BASE_URL/health")
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 成功: $response${NC}"
else
    echo -e "${RED}✗ 失败${NC}"
fi
echo ""

# 2. 用户注册
echo -e "${YELLOW}2. 用户注册...${NC}"
response=$(curl -s -X POST "$BASE_URL/auth/register" \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "email": "test@example.com",
        "password": "password123",
        "full_name": "Test User"
    }')

if echo "$response" | grep -q '"success":true'; then
    username=$(echo "$response" | grep -o '"username":"[^"]*"' | cut -d'"' -f4)
    echo -e "${GREEN}✓ 成功: 用户 $username 已创建${NC}"
else
    echo -e "${RED}✗ 失败: $response${NC}"
fi
echo ""

# 3. 用户登录
echo -e "${YELLOW}3. 用户登录...${NC}"
response=$(curl -s -X POST "$BASE_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "password": "password123"
    }')

if echo "$response" | grep -q '"success":true'; then
    TOKEN=$(echo "$response" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    echo -e "${GREEN}✓ 成功: 已获取 Token${NC}"
    echo -e "${CYAN}Token: $TOKEN${NC}"
else
    echo -e "${RED}✗ 失败: $response${NC}"
fi
echo ""

# 检查 token 是否获取成功
if [ -z "$TOKEN" ]; then
    echo -e "${RED}错误: 无法获取 Token，后续需要认证的测试将跳过${NC}"
    exit 1
fi

# 4. 获取当前用户信息
echo -e "${YELLOW}4. 获取当前用户信息...${NC}"
response=$(curl -s -X GET "$BASE_URL/me" \
    -H "Authorization: Bearer $TOKEN")

if echo "$response" | grep -q '"success":true'; then
    username=$(echo "$response" | grep -o '"username":"[^"]*"' | cut -d'"' -f4)
    echo -e "${GREEN}✓ 成功: 用户 $username${NC}"
else
    echo -e "${RED}✗ 失败: $response${NC}"
fi
echo ""

# 5. 更新当前用户信息
echo -e "${YELLOW}5. 更新当前用户信息...${NC}"
response=$(curl -s -X PUT "$BASE_URL/me" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
        "full_name": "Updated Test User",
        "email": "updated@example.com"
    }')

if echo "$response" | grep -q '"success":true'; then
    echo -e "${GREEN}✓ 成功: 用户信息已更新${NC}"
else
    echo -e "${RED}✗ 失败: $response${NC}"
fi
echo ""

echo -e "${CYAN}=====================================${NC}"
echo -e "${CYAN}测试完成${NC}"
echo -e "${CYAN}=====================================${NC}"
